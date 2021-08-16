package stats
// reference：https://github.com/akhenakh/statgo

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L${SRCDIR}/lib -lstatgrab
// #include "statgrab.h"
import "C"
import (
	"runtime"
	"sync"
)

// Stat handle to access libstatgrab
type Stat struct {
	sync.Mutex
	exitMessage chan bool
}

// NewStat return a new Stat handle
func NewStat() *Stat {
	s := &Stat{}
	runtime.SetFinalizer(s, (*Stat).Free)

	initDone := make(chan bool)
	s.exitMessage = make(chan bool)

	C.sg_init(1)

	go func() {
		// We need some function calls to be performed on the same thread
		// Those for which statgrab is using a thread local
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		// Throw away the first reading as thats averaged over the machines uptime
		C.sg_get_cpu_stats_diff(nil)
		C.sg_get_network_io_stats_diff(nil)
		C.sg_get_page_stats_diff(nil)
		C.sg_get_disk_io_stats_diff(nil)

		initDone <- true

		for {
			select {
			case <-s.exitMessage:
				return
			case f := <-mainfunc:
				f()
			}
		}

	}()

	<-initDone

	return s
}

func (s *Stat) Free() {
	s.Lock()
	C.sg_shutdown()
	s.exitMessage <- true
	s.Unlock()
}

// queue of work to run in main thread.
var mainfunc = make(chan func())

// do runs f on the main thread.
func do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}
