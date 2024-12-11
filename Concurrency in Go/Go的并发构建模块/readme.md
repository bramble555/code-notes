# goroutine

在 go1.23里面

```go
func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}

	wg.Wait()
}


```

和下面这个程序运行结果一样

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {  // 显式传递 salutation
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // 将 salutation 作为参数传递
	}

	wg.Wait()
}

```

似乎在 go1.18里面 go底层优化了,在 1.18以前, 第一个程序会打印出来

``` 
good day  good day  good day 
```

了解一下 goroutine 历史

## Go 栈大小调整历史

1. **Go 1.0 到 Go 1.1**：`goroutine` 的栈最初是 **4 KB**。
2. **Go 1.2**：将 `goroutine` 初始栈大小调整为 **8 KB**。
3. **Go 1.3**：引入了 **连续栈**，将 `goroutine` 的初始栈大小降低到 **2 KB**，并且栈可以在执行过程中动态扩展和收缩。
4. **Go 1.4**：保持 `goroutine` 的栈初始大小为 **2 KB**，最大栈大小为 **1 GB**。
5. **Go 1.5 及以后**：`goroutine` 栈的初始大小仍为 **2 KB**，但是随着栈增长会动态扩展，栈空间大小会根据实际需求调整。

# Sync package

## WaitGroup

技巧 wg.Done 在 并发函数内部

什么意思 ?

Example

```go
go func() {
			defer arithmetic.Done()
			decrement()
		}()
increment := func() {
    lock.Lock() // 1
    defer lock.Unlock() // 2
    count++
    fmt.Printf("Incrementing: %d\n", count)
}

或者	

go func() {
    defer wg.Done() //2
    fmt.Println("1st goroutine sleeping...")
}()		

```

如果在 `goroutine` 函数外部调用 wg.Done(),一定会死锁.为什么?因为你可能函数已经调用了,但是 wg 没 -1

## RWMutex 和 Mutex

读写操作

要求:只要没有别的东西占用写操作，任意数量的读取者就可以进行读取操作

所以此时 sync.RWMutex 比 sync.Mutex  性能在并发量很高的时候性能较好

补充嘿嘿没见过的东西

tabwriter模块  可以对齐输出



## Cond

Cond实现了一个条件变量，用于等待或宣布事件发生时goroutine的交汇点。

`sync.Cond` 的设计目的是为了协调多个 goroutine 的运行. 让 多个 goroutine  顺序执行



像这样

```go
c := sync.NewCond(&sync.Mutex{}) // 1
c.L.Lock() // 2
for conditionTrue() == false {
    c.Wait() // 3
}
c.L.Unlock() // 4

当然后面需要唤醒
```

或者

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	conditionMet := false // 用于表示条件是否满足

	wg.Add(2) // 需要等待两个 goroutine 完成

	// Goroutine 1: 等待条件满足
	go func() {
		defer wg.Done() // 标记当前 goroutine 完成
		mu.Lock()
		// 如果 conditionMet 为 false，那么循环会继续；如果 conditionMet 为 true，循环会退出。
		for !conditionMet { // 使用 for 循环检查条件，防止虚假唤醒
			fmt.Println("Goroutine 1: 等待条件满足...")
			cond.Wait() // 释放 mu 锁，进入等待
		}
		fmt.Println("Goroutine 1: 条件满足，继续执行")
		mu.Unlock()
	}()

	// Goroutine 2: 模拟条件发生变化
	go func() {
		defer wg.Done() // 标记当前 goroutine 完成
		mu.Lock()
		fmt.Println("Goroutine 2: 条件已满足，发送信号")
		conditionMet = true // 更新条件状态
		cond.Signal()       // 唤醒一个等待中的 goroutine
		mu.Unlock()
	}()

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("所有 goroutine 执行完毕")
}

```

Notes:

​	**`c.Wait()`**：它会释放 `Cond.L` 锁并将当前 goroutine 阻塞，直到被另一个 goroutine 唤醒（通常是通过调用 `Signal()` 或 `Broadcast()`）。  区分 wg.Wait()



我的理解:  cond.Wait() 会导致上下文切换(除非你只有一个goroutine ),那么相当于给目前的 goroutine 上了锁,然后切换到另外一个 goroutine ,需要另外一个 goroutine 解锁!  当然 ,这是方便理解,而不是真正的上锁了



来个题吧

假设我们有一个固定长度为2的队列，并且我们要将10个元素放入队列中。 我们希望一有空间就能放入，所以在队列中有空间时需要立刻通知.

如何放?  10个元素怎么放入 长度为2的队列?  一直放,当长度为2 的时候,取走,长度小于2,再放  懂了吧

那么这里的 长度为2 的时候就要等  wait, 等取走,取走完了 给你信号,你可以继续生产了,生产到 长度为2的时候 继续等

```go
func main() {
	var queue = make([]int, 0, 10)
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	// 使用 WaitGroup 确保主函数等待 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(1)
	// 启动消费者
	go func() {
		defer wg.Done()
		for {
			cond.L.Lock()
			// 要用 for 循环进行等待，否则可能假唤醒
			for len(queue) == 0 {
				// 等待生产者
				cond.Wait() // tag:1
			}
			fmt.Println("我消费了", queue[0])
			if queue[0] == 9 {
				queue = queue[1:]
				cond.Signal()
				cond.L.Unlock()
				break
			}
			queue = queue[1:]
			cond.Signal() // 对应 tag2

			cond.L.Unlock()

		}

	}()

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		// 要用 for 循环进行等待，否则可能假唤醒
		for len(queue) == 2 {
			// 切换到另外一个 goroutine,那个 goroutine 执行完 需要唤醒我这个 goroutine
			// 等待，切换到消费者
			cond.Wait() // tag:2
		}

		queue = append(queue, i)
		// 不要用 defer,因为这不是一个函数,只是在 for 循环里面
		cond.Signal() // 对应 tag1
		cond.L.Unlock()

	}
	wg.Wait()
}
```

这里面 有很多点需要注意

- cond.Wait() 和 cond.Signal() 一一对应

- 消费者退出条件

  ```go
  if queue[0] == 9 {
  			queue = queue[1:]
  			cond.Signal()
  			cond.L.Unlock()
  			break
  		}
  如果不需要切换到 另外一个goroutine 可以不用+cond.Signal()
  ```
  如果生产者还要干其他事情, 那么可能就要 wg.Add(2)

- ```
  for len(queue) == 0 {
  				// 等待生产者
  				cond.Wait() // tag:1
  			}
  ```

  一定要用 for 循环进行等待,否则可能出现假等待

- 如果不使用 var wg sync.WaitGroup,消费者提前退出, wg.Done() 保证确实消费到 9 了.

- defer 是在函数末尾退出,如果你在 for 循环里面用,相当于 在 main 函数最后用了 defer,所以不要这样做.

  当然,我的代码不够简洁,还有更加简洁的版本.但是这个方便理解.



cond.Broadcast() 能 唤醒所有等待的 goroutine

书上的代码场景有些复杂

用到的时候再来看

## Once

```
func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(decrement)
	once.Do(increment)

	fmt.Printf("Count: %d\n", count)
}

答案是1

func main() {
	var count int

	// 定义两个函数
	increment := func() { count++ }
	decrement := func() { count-- }

	var onceIncrement sync.Once
	var onceDecrement sync.Once

	// 每个函数都通过各自的 sync.Once 来确保只执行一次
	onceIncrement.Do(increment)
	onceDecrement.Do(decrement)

	// 打印 count 的值
	fmt.Printf("Count: %d\n", count)
}

答案是0
```

var onceIncrement sync.Once  只能结构体名字所示,任何函数,只执行一次.

如果想执行多次,你就要定义多个结构体来执行不同的函数

## Pool

在较高的层次上，池模式是一种创建和提供固定数量可用对象的方式。它通常用于约束创建资源昂贵的事物（例如数据库连接）。Go的sync.Pool可以被多个例程安全地使用。

Pool的主要接口是它的Get方法。 被调用时，Get将首先检查池中是否有可用实例返回给调用者，如果没有，则创建一个新成员变量。使用完成后，调用者调用Put将正在使用的实例放回池中供其他进程使用。



```go
func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem // 1
		},
	}

	// 将池扩充到 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // 2
			defer calcPool.Put(mem)

		}()
	}

	// 假设内存中执行了一些快速的操作

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
```

4KB要 进行分配给 1MB, 假设打印 是10calculators were created. 那么说明在这个过程 又只创建了 6次 1KB.



如果使用池子里东西在内存上不是大致均匀的，则会花更多时间将从池中检索，这比首先实例化它要耗费更多的资源。例如，你的程序需要随机和可变长度的切片，在这种情况下Pool不会为你提供太多的帮助。

什么意思?  就是 取出和放回的东西不一样,就别用 Pool 了

因此，在使用Pool时，请记住以下几点：

- 实例化sync.Pool时，给它一个新元素，该元素应该是线程安全的。
- 当你从Get获得一个实例时，不要假设你接收到的对象状态。(对象可能未初始化,在使用对象之前，最好检查或初始化对象的状态，以确保它符合你的预期)
- 当你从池中取得实例时，请务必不要忘记调用Put。否则池的优越性就体现不出来了。这通常用defer来执行延迟操作。
- 池中的元素必须大致上是均匀的。



# Channel



试将值写入只读通道或从只写通道读取值都是错误的。

```go
writeStream := make(chan<- interface{})
readStream := make(<-chan interface{})

<-writeStream
readStream <- struct{}{}
```

千万不要按照上面的写

下面才是正确的!!

```go

	dataStream := make(chan interface{})

	// 声明只写通道和只读通道
	var writeStream chan<- interface{} // 只写通道
	var readStream <-chan interface{}  // 只读通道

	// 将双向通道赋值给只写通道和只读通道
	writeStream = dataStream
	readStream = dataStream
```

< - 运算符的接收形式也可以选择返回两个值

```
stringStream := make(chan string)
go func() {
    stringStream <- "Hello channels!"
}()
salutation, ok := <-stringStream // 1
fmt.Printf("(%v): %v", ok, salutation)

当管道没有关闭的时候,ok 是true,否则为 false
```

管道能关闭!!

这为我们开辟了一些新的模式。

1. 通道的range操作

   `range` 循环会一直从通道中接收数据，直到通道被关闭并且没有剩余数据.我之前一直理解是错的, `range` 不是必须等到通道关闭才能开始接收数据。我现在知道了  close 之后,range 才能取走数据,直到管道里面没有数据.

   综上:`range` 可以在通道关闭之前就开始接收数据，只要通道有数据可读。当通道关闭时，`range` 会继续读取直到没有数据可读取为止。

```
intStream := make(chan int)
go func() {
    defer close(intStream) // 1
    for i := 1; i <= 5; i++ {
        intStream <- i
    }
}()

for integer := range intStream { // 2
    fmt.Printf("%v ", integer)
}
```

​	奥对 ,这里细节, intStream 是**无缓冲通道**，它每次发送一个数据，必须等待接收方接收才能继续发送。注意循环退出并没有设置条件，并且range也不返回第二个布尔值

2. 关闭通道后，多个 goroutine 可以同时停止等待

   假设你有多个 goroutine，它们都在等待从一个通道中接收数据。你可以通过关闭通道来通知它们不再继续等待，而不需要逐个通知每个 goroutine
   
   所以 关闭某个通道同样可以被当作向多个goroutine同时发生消息的方式之一

Example

```go
func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin //1
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	time.Sleep(time.Second * 10)
	close(begin) //2
	wg.Wait()
}
```

这个例子  打印 Unblocking goroutines...  然后睡了10秒,打印

3 has begun  0 has begun   2 has begu ....   后面还有 1和 4

当然 也可以用 sync.Cond  Single或者Brocast  来替代

```
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			mu.Lock() // 锁住互斥锁
			cond.Wait() // 等待条件满足
			mu.Unlock() // 释放锁

			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	time.Sleep(time.Second * 10)

	// 使用 Signal 逐个唤醒 goroutine
	for i := 0; i < 5; i++ {
		mu.Lock()
		cond.Signal() // 唤醒一个等待的 goroutine
		mu.Unlock()
	}

	wg.Wait()
}

```

我们说向一个已满的通道写入，会出现阻塞，从一个已空的通道读取，也会出现阻塞。这里的“满”和“空”是针对容量或缓冲区大小而言的。

管道很容易产生 block 和 panic,正如表格所示(表格 是在 main函数执行的)

![img](https://www.topgoer.cn/uploads/concurrency/images/m_43befaa5760c9dc57030745c911a17ec_r.jpg)

因此强烈建议你在自己的程序中尽可能做到保持通道覆盖范围最小，以便这些事情保持明显。如果你将一个通道作为一个结构体的成员变量，并且有很多方法，它很快就会把你自己给绕进去

保持通道覆盖范围最小  像这样

```
chanOwner := func() <-chan int {

    resultStream := make(chan int, 5)//1
    go func() {//2
        defer close(resultStream)//3
        for i := 0; i <= 5; i++ {
            resultStream <- i
        }
    }()
    return resultStream//4

}

resultStream := chanOwner()
for result := range resultStream {//5
    fmt.Printf("Received: %d\n", result)
}
fmt.Println("Done receiving!")
```



来做个 作业:   三个工人干五个活,最快效率完成

```
const worker int = 3
const task int = 5

func main() {
	var wg sync.WaitGroup
	var work func(id int, channelWork <-chan int, wg *sync.WaitGroup)
	work = func(id int, channelWork <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := range channelWork {
			log.Printf("%d工人正在干%d活", id, i)
			time.Sleep(time.Second * 2)
		}
	}
	channelWork := make(chan int, 2)

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go work(i, channelWork, &wg)
	}
	for i := 0; i < task; i++ {
		channelWork <- i
	}

	close(channelWork)
	wg.Wait()
	fmt.Println("干完了")

}
```

缓冲区大小多少都可以.

管道要有很多细节注意,比如 

- 初始化管道
- 关闭管道(不要多次关闭)
- 不要读和写  关闭的管道

# Select

select会一直等待直到某个case语句完成

```go
start := time.Now()
c := make(chan interface{})
go func() {
    time.Sleep(5 * time.Second)
    close(c) // 1
}()

fmt.Println("Blocking on read...")
select {
case <-c: // 2
    fmt.Printf("Unblocked %v later.\n", time.Since(start))
}

输出
Blocking on read...
Unblocked 5s later.
```

这段代码能正常运行的原因是，`channel` 本身并不需要立即有数据发送，关闭 `channel` 也能在 `select` 语句中触发阻塞解除

再来看一个例子

```go
func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 999; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

输出不一定是啥
c1Count: 505
c2Count: 495
```

Go运行时对一组case语句执行伪随机统一选择。这意味着在同样的条件下，每个case被选中的机会几乎是一样的。



如果所有通道都尚未初始化完成会发生什么？如果所有的通道都处在阻塞状态，你无法进行处理，但你又不能就这样持续阻塞下去，你可能希望程序能够执行超时。Go的time包提供了一个很好的方式来完成这个功能，这些功能完全符合选择语句的范式。 这里有一个例子：

```go
var c <-chan int
select {
case <-c: //1
case <-time.After(1 * time.Second):
    fmt.Println("Timed out.")
}
输出
Timed out.
```

如果我们想做点什么，但当前通道还没准备好呢？select语句中允许我们添加default条件，以便你在所有分支都不符合执行条件的时候执行。

`select` 会 **阻塞等待** 直到有可用的信号，或者所有的 `case` 都不满足时，才会执行 `default`

```go
start := time.Now()
var c1, c2 <-chan int
select {
case <-c1:
case <-c2:
default:
    fmt.Printf("In default after %v\n\n", time.Since(start))
}		
输出
In default after 0s
```

可以看到它几乎是瞬间运行默认语句。这允许你在不阻塞的情况下退出选择块

上面这个例子 执行完 select 立马就退出了,

那么我们现在想 执行 select的时候 不退出,当某个条件达成后再退出, 那就用到了 for-select-case

```go
done := make(chan interface{})
go func() {
    time.Sleep(5 * time.Second)
    close(done)
}()

workCounter := 0
loop:
for {
    select {
    case <-done:
        break loop
    default:
    }

    // Simulate work
    workCounter++
    time.Sleep(1 * time.Second)
}

fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)

打印
Achieved 5 cycles of work before signalled to stop.
```

细节:

1. default 没做任何事情,只是为了 done 没有接收到信号的时候, workCounter++ 能够执行,

2. ```
   workCounter++
       time.Sleep(1 * time.Second)
   ```

   把这个代码片段 放入 default 里面 打印内容也是相同的



最后一个知识点

```
select {}
```

这个语句将永远阻塞

# GOMAXPROCS

PROCS: 是“处理器”（processors）的缩写，在这里指的是可以并行执行goroutine的逻辑处理器

1.5 之后 默认设置主机上的逻辑CPU数量,以前总是1

如果你想调整这个值呢？ 大多数时候你尽量别这么想。Go的调度算法在大多数情况下足够好，即增加或减少工作队列和线程的数量可能会造成更多的伤害，但仍然有些情况下可能会更改此值。

出问题了,你再尝试改.



ok.文章就结束了,你已经学会了 何时用锁,何时使用通道和select语句来 "通过通信共享内存".