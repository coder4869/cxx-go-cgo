## goroutine
日常开发中，经常会使用后台 goroutine 做一些定期清理或者更新的任务。

### 生命周期
goroutine 生命周期基本和主程序一致。

goroutine 关闭方式：

- **自动回收：SetFinalizer**
    * 特点：不需要用户显式关闭，在检查到没有引用之后，主动终止 goroutine，等待 gc 回收。
    * 声明：`func SetFinalizer(obj interface{}, finalizer interface{})`
    * 用法：`runtime.SetFinalizer(obj, (*StructType).free)` StructType 是 obj 的数据类型，free 是 StructType 的一个成员函数方法。
    * 机制：对象可以关联一个 SetFinalizer 函数，当 gc 检测到 unreachable 对象有关联的 SetFinalizer 函数时，会执行关联的 SetFinalizer 函数，同时取消关联。在下一次 GC时，对象重新处于 unreachable 状态并且没有 SetFinalizer 关联，就会被回收。
    * 注意事项
        * 不论是程序正常结束、亦或发生错误，对象被 gc 选中并被回收之前，SetFinalizer 都不会执行，因此不要在 SetFinalizer 中执行将内存中的内容 flush 到磁盘之类的操作。
        * SetFinalizer 延长了对象的生命周期：在第一次 GC 时执行 Finalizer 函数，且目标对象重新变成可达状态，直到第二次才真正 “销毁”。这对有大量对象分配的高并发算法，可能会造成很大麻烦。
        * 指针构成的 "循环引⽤" 加上 runtime.SetFinalizer 会导致内存泄露