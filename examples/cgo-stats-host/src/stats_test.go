package stats_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"runtime"
	"sync"
	"testing"
	"time"
	"go_stats/stats"

	// "github.com/stretchr/testify/assert"
)

func TestHostInfo(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	hi := s.HostInfos()
	// assert.NotNil(t, s)
	// assert.NotEmpty(t, hi.HostName, hi.OSName, hi.OSRelease, hi.OSVersion, hi.Platform)
	// assert.True(t, hi.NCPUs > 0, hi.MaxCPUs > 0)

	pst := time.Now().UnixNano()
	fmt.Println("\nHostInfos Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print hi by calling HostInfos.String()
	t.Log(hi)
}

func TestCPU(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	cpu := s.CPUStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, cpu)

	time.Sleep(20 * time.Millisecond)

	cpu = s.CPUStats()

	pst := time.Now().UnixNano()
	fmt.Println("\nCPUStats Time Cost: ", (pst - st) / int64(time.Millisecond) - 20)
	// print cpu by calling CPUStats.String()
	t.Log(cpu)
}

func TestCPULoad(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	// initialCPU := s.CPUStats()
	s.CPUStats()
	cpuCount := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuCount)
	var wg sync.WaitGroup
	wg.Add(cpuCount)
	doneChan := make(chan bool, cpuCount)
	for k := 0; k < cpuCount; k++ {
		go func() {
			defer wg.Done()
			var i uint64 = 2
			for {
				select {
				case <-doneChan:
					{
						return
					}
				default:
					{
					}
				}
				i = i * i
			}
		}()
	}
	s.CPUStats()
	testDuration := 5 * time.Second
	time.Sleep(testDuration)
	cpu := s.CPUStats()
	for k := 0; k < cpuCount; k++ {
		doneChan <- true
	}
	wg.Wait()

	pst := time.Now().UnixNano()
	fmt.Println("\nTestCPULoad Time Cost: ", (pst - st) / int64(time.Millisecond) - 5 * 1000)
	// print cpu by calling CPUStats.String()
	t.Log(cpu)

	// //Assure that the Period of the stats is about the same as the Duration of the stats.
	// assert.True(t, cpu.Period-testDuration < time.Second*2 || cpu.Period-testDuration > 2*time.Second)
	// t.Logf("CPU Idle %f:", cpu.Idle)
	// //The CPU should not be idle if we run cpuCount goroutines
	// assert.True(t, cpu.Idle < 50.0)
	// //The stats should have changed from the start till the finish of this test
	// assert.True(t, cpu.Idle != initialCPU.Idle)
}

func TestFSInfos(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	f := s.FSInfos()
	// assert.True(t, len(f) > 0)

	pst := time.Now().UnixNano()
	fmt.Println("\nFSInfos Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print fs by calling FSInfos.String()
	for _, fs := range f {
		t.Log(fs)
	}
}

func TestInterfaceInfos(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	interfaces := s.InteraceInfos()
	// assert.True(t, len(interfaces) > 0)

	pst := time.Now().UnixNano()
	fmt.Println("\nInteraceInfos Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print i by calling InteraceInfos.String()
	for _, i := range interfaces {
		t.Log(i)
	}
}

func TestVM(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	m := s.MemStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, m)

	pst := time.Now().UnixNano()
	fmt.Println("\nMemStats Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print m by calling MemStats.String()
	t.Log(m)
}

func TestDisksIO(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	d := s.DiskIOStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, d)

	pst := time.Now().UnixNano()
	fmt.Println("\nDiskIOStats Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print d by calling DiskIOStats.String()
	t.Log(d)
}

func TestNetIO(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	n := s.NetIOStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, n)

	pst := time.Now().UnixNano()
	fmt.Println("\nNetIOStats Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print n by calling NetIOStats.String()
	t.Log(n)
}

func getInterface(arr []*stats.NetIOStats, names ...string) (*stats.NetIOStats, error) {
	for _, ns := range arr {
		for _, name := range names {
			if ns.IntName == name {
				return ns, nil
			}
		}
	}
	return nil, fmt.Errorf("No interface not found matching any of %v", names)
}

func TestNetIOTXRX(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	beforeNetIOArr := s.NetIOStats()
	beforeNetIO, err := getInterface(beforeNetIOArr, "lo", "lo0")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Log("Could not listen on port 8080", err)
		t.SkipNow()
		return
	}
	defer ln.Close()
	go func() {
		conn, aerr := ln.Accept()
		if aerr != nil {
			return
		}
		defer conn.Close()
		io.Copy(ioutil.Discard, conn)
	}()
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Log("Could not connect to server:", err)
		t.SkipNow()
		return
	}
	defer conn.Close()
	conn.Write(make([]byte, 1024*1024))
	afterNetIOArr := s.NetIOStats()
	afterNetIO, err := getInterface(afterNetIOArr, "lo", "lo0")
	if err != nil {
		t.Log(err)
		t.SkipNow()
		return
	}

	pst := time.Now().UnixNano()
	fmt.Println("\nTestNetIOTXRX Time Cost: ", (pst - st) / int64(time.Millisecond))

	t.Log("BeforeData:", beforeNetIO)
	t.Log("After 1MB data:", afterNetIO)
	// assert.True(t, beforeNetIO.TX < afterNetIO.TX)
	// assert.True(t, beforeNetIO.RX < afterNetIO.TX)
}

func TestProcess(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	p := s.ProcessStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, p)

	pst := time.Now().UnixNano()
	fmt.Println("\nProcessStats Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print p by calling ProcessStats.String()
	t.Log(p)
}

func TestPages(t *testing.T) {
	st := time.Now().UnixNano()

	s := stats.NewStat()
	p := s.PageStats()
	// assert.NotNil(t, s)
	// assert.NotNil(t, p)

	pst := time.Now().UnixNano()
	fmt.Println("\nPageStats Time Cost: ", (pst - st) / int64(time.Millisecond))
	// print p by calling PageStats.String()
	t.Log(p)
}

func TestGoRoutineCleanup(t *testing.T) {
	st := time.Now().UnixNano()

	var wg sync.WaitGroup

	s := stats.NewStat()
	s.PageStats()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Free()
		pst := time.Now().UnixNano()
		fmt.Println("\nTestGoRoutineCleanup Time Cost: ", (pst - st) / int64(time.Millisecond))
	}()

	wg.Wait()
	pst := time.Now().UnixNano()
	fmt.Println("\nTestGoRoutineCleanup wg.Wait() Time: ", (pst - st) / int64(time.Millisecond))
}

func TestGoRoutines(t *testing.T) {
	// test for ticket #2
	// ping -s 20000 localhost, check for growing lo0 stats ([0] at least on OSX)

	t.Skip()
	var wg sync.WaitGroup
	s := stats.NewStat()

	wg.Add(1)
	go func() {
		defer wg.Done()
		t.Log("1", s.NetIOStats()[0])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2500 * time.Millisecond)
		t.Log("2", s.NetIOStats()[0])
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		t.Log("4", s.NetIOStats()[0])
	}()
	wg.Wait()
}
