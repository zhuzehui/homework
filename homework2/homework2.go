package homework

import (
	"sync"
	"sync/atomic"
	"time"
)

func main() {

}

// 定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10
func increaseByTen(num *int) {
	*num += 10
}

// 接收一个整数切片的指针，将切片中的每个元素乘以2
func multiplyByTwo(slice *[]int) {
	for i := range *slice {
		(*slice)[i] *= 2
	}
}

// 使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOddEven() {
	// 创建一个通道用于同步
	done := make(chan bool)

	// 协程打印奇数
	go func() {
		for i := 1; i <= 10; i += 2 {
			println(i)
		}
		done <- true // 通知主协程完成
	}()

	// 协程打印偶数
	go func() {
		for i := 2; i <= 10; i += 2 {
			println(i)
		}
		done <- true // 通知主协程完成
	}()

	// 等待两个协程完成
	<-done
	<-done
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func taskScheduler(tasks []func()) {
	done := make(chan bool)
	for _, task := range tasks {
		go func(t func()) {
			start := time.Now()
			t()
			end := time.Now()
			println("Task completed in:", end.Sub(start))
			done <- true
		}(task)
	}
	for i := 0; i < len(tasks); i++ {
		<-done
	}
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	println("Name:", e.Name)
	println("Age:", e.Age)
	println("Employee ID:", e.EmployeeID)
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func channelCommunication() {
	ch := make(chan int)

	// 生成整数的协程
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // 关闭通道，表示没有更多的值发送
	}()

	// 接收整数的协程
	go func() {
		for num := range ch {
			println(num)
		}
	}()
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func bufferedChannel() {
	ch := make(chan int, 10) // 创建一个带缓冲的通道，缓冲区大小为10

	// 生产者协程
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch) // 关闭通道，表示没有更多的值发送
	}()

	// 消费者协程
	go func() {
		for num := range ch {
			println(num)
		}
	}()
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func mutexCounter() {
	var counter int
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(time.Second) // 等待所有协程完成
	println("Counter value:", counter)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atomicCounter() {
	var counter int64

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	time.Sleep(time.Second) // 等待所有协程完成
	println("Counter value:", counter)
}
