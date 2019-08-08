package concurrency

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

/**
因为mu.Lock()和mu.Unlock()并不在同一个Goroutine中，所以也就不满足顺序一致性内存模型。
同时它们也没有其它的同步事件可以参考，这两个事件不可排序也就是可以并发的。
因为可能是并发的事件，所以main函数中的mu.Unlock()很有可能先发生，而这个时刻mu互斥对象还处于未加锁的状态，从而会导致运行时异常。
*/
func TestConcurrency(t *testing.T) {
	var mu sync.Mutex
	go func() {
		fmt.Println("Hello World")
		mu.Lock()
	}()
	mu.Unlock()
}

/*
修复的方式是在main函数所在线程中执行两次mu.Lock()，当第二次加锁时会因为锁已经被占用（不是递归锁）而阻塞，
main函数的阻塞状态驱动后台线程继续向前执行。当后台线程执行到mu.Unlock()时解锁，此时打印工作已经完成了，
解锁会导致main函数中的第二个mu.Lock()阻塞状态取消，此时后台线程和主线程再没有其它的同步事件参考，
它们退出的事件将是并发的：在main函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。
虽然无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
 *
*/
func TestCon1(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("Hello World")
		mu.Unlock()
	}()
	mu.Lock()
}

/**
使用sync.Mutex互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步：
根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成（从而退出main、退出程序），而此时打印工作已经完成了。
如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了
*/
func TestChanSync(t *testing.T) {
	done := make(chan int) //构建无缓冲channel , **对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
	go func() {
		fmt.Println("Hello World")
		<-done
	}()
	done <- 1
}

func TestChanSync1(t *testing.T) {
	done := make(chan int, 2) //构建缓冲channel
	go func() {
		fmt.Println("Hello World")
		done <- 1
	}()
	<-done
}

func TestChanSync2(t *testing.T) {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

/**
等待待N个线程完成后再进行下一步的同步操作有一个简单的做法，就是使用sync.WaitGroup来等待一组事件
*/
func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello World")
			wg.Done()
		}()
	}
	wg.Wait()
}

/**
生产者 消费者模式
*/
func TestProCon(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(1, ch)
	go Consumer(ch, "小张")
	go Consumer(ch, "小李")
	time.Sleep(1 * time.Millisecond)
}

/**
上面的function 靠休眠方式是无法保证稳定的输出结果的。
我们可以让main函数保存阻塞状态不退出，只有当用户输入Ctrl-C时才真正退出程序
*/
func TestProCon1(t *testing.T) {
	ch := make(chan int, 64)
	go Producer(1, ch)
	go Consumer(ch, "小张")
	go Consumer(ch, "小李")
	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}

// 生产者: 生成 factor 整数倍的序列
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		fmt.Println("生产数据：", i*factor)
		out <- i * factor
	}
}

func Consumer(in <-chan int, name string) {
	for v := range in {
		fmt.Println(name, "----消费数据：", v)
	}
}
