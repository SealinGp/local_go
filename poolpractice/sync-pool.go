package poolpractice

import (
	"bytes"
	"log"
	"net"
	"sync"

	"github.com/fatih/pool"
	"github.com/valyala/bytebufferpool"
)

/**
取出来的 bytes.Buffer 在使用的时候，我们可以往这个元素中增加大量的 byte 数据，
这会导致底层的 byte slice 的容量可能会变得很大。这个时候，即使 Reset 再放回
到池子中，这些 byte slice 的容量不会改变，所占的空间依然很大。而且，因为 Pool
回收的机制，这些大的 Buffer 可能不被回收，而是会一直占用很大的空间，这属于内存泄漏的问题,
所以:
1.在回收超过一定大小容量(64KB)的buffer时, 可以直接丢弃，不用回收了
2.可以将池子分成几个量级的池子,根据需要在合适的大小池子中获取了, 如 0~512Byte 512Byte~1KB 1KB~4KB
1MB = 1024 KB
1KB = 1024 B

32*1024 Byte = 32 KB
*/

type buf1 struct {
	pool sync.Pool
}

func NewBuf1() *buf1 {
	buf := &buf1{
		pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(nil)
			},
		},
	}
	return buf
}

func (buf *buf1) GetBuffer() *bytes.Buffer {
	return buf.pool.Get().(*bytes.Buffer)
}

func (buf *buf1) PutBuffer(bu *bytes.Buffer) {
	maxSize := 1 << 16
	if bu.Len() > maxSize {
		return
	}

	bu.Reset()
	buf.pool.Put(bu)
}

func TestBuf() {
	buf1 := NewBuf1()

	f := buf1.GetBuffer()

	f.Write([]byte("123"))
	log.Printf("[E] %v", f.String())

	buf1.PutBuffer(f)
}

func TestBuf1() {
	byteBuffer := bytebufferpool.Get()
	byteBuffer.Write([]byte("first line \n"))
	byteBuffer.Write([]byte("second line \n"))
	byteBuffer.B = append(byteBuffer.B, "third line \n"...)

	log.Printf("[I] byebuff contents=%q", byteBuffer.Bytes())

	bytebufferpool.Put(byteBuffer)
}

//连接池
func TestTcpPool() {
	factory := func() (net.Conn, error) {
		return net.Dial("tcp", "127.0.0.1:4000")
	}

	p, err := pool.NewChannelPool(5, 30, factory)

	conn, err := p.Get()
	if err != nil {
		log.Printf("[E] get connect failed. err:%v", err)
		return
	}

	conn.Close()

	if pc, ok := conn.(*pool.PoolConn); ok {
		pc.MarkUnusable()
		pc.Close()
	}

	p.Close()

	cuurent := p.Len()
}

func TestWorkerPool() {

}
