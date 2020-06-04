package timewheel

import (
	"container/list"
	"sync"
	"sync/atomic"
	"unsafe"
)

//定时器任务
type Timer struct {
	expiration int64 //延迟多少milliseconds后执行
	task func()      //待执行的任务

	//该定时器属于哪个槽,该槽的指针
	//此字段可能会通过Timer.Stop()和Bucket.Flush() 同时更新和读取
	//该指针类型为*bucket
	b unsafe.Pointer

	element *list.Element //链表节点
}

func (t *Timer)getBucket() *bucket {
	return (*bucket)(atomic.LoadPointer(&t.b))
}
func (t *Timer)setBucket(b *bucket)  {
	atomic.StorePointer(&t.b,unsafe.Pointer(b))
}

func (t *Timer)Stop() bool {
	stopped := false
	for b := t.getBucket(); b != nil; b = t.getBucket() {
		stopped = b.Remove(t)
	}
	return stopped
}

//槽
type bucket struct {
	//64位(bit)原子性操作,32位不确定,所以我们必须让64位的字段作为结构体的第一个字段
	expiration int64  //延迟多少milliseconds后执行

	mu         sync.Mutex
	timers    *list.List   //槽里面对应的定时器链表
}

func newBucket() *bucket {
	return &bucket{
		timers:list.New(),
		expiration:-1,
	}
}

//原子性读取数据
func (b *bucket)Expiration() int64 {
	return atomic.LoadInt64(&b.expiration)
}
//原子性写入数据
func (b *bucket)SetExpiration(new int64) bool {
	return atomic.SwapInt64(&b.expiration,new) != new
}

//添加定时器
func (b *bucket)Add(t *Timer)  {
	b.mu.Lock()

	newEle := b.timers.PushBack(t)
	t.setBucket(b)
	t.element = newEle

	b.mu.Unlock()
}

func (b *bucket)remove(t *Timer) bool {
	if t.getBucket() != b {
		//如果remove函数是被t.Stop调用,并且发生在以下情况:
		//1.从b里面移除t (通过b.Flush 调用b.remove)时
		//2.吧t从b槽移动到another bucket(另外一个槽)时,(b.Flush 调用b.remove和 ab.Add(ab = another bucket = 另外一个时间轮上的槽))
		//情况1:getBucket将会返回 nil
		//情况2:getBucket将会返回 非nil(返回ab=another bucket)
		//在这两种情况下,该bucket != 当前的bucket,因此无需进行移除操作
		return false
	}

	b.timers.Remove(t.element)
	t.setBucket(nil)
	t.element = nil
	return true
}

func (b *bucket)Remove(t *Timer) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.remove(t)
}

//删除当前槽
func (b *bucket)Flush(reinsert func(*Timer))  {
	b.mu.Lock()
	e := b.timers.Front()
	for e != nil {
		next := e.Next()
		t    := e.Value.(*Timer)
		b.remove(t)
		reinsert(t)
		e = next
	}
	b.mu.Unlock()

	b.SetExpiration(-1)
}