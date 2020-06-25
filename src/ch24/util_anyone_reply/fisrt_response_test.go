package concurrency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string)
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			// 放消息，如果不是 buffer channel
			// return 之后 没有从 channel 中取消息的 receiver
			// 导致协程阻塞
			ch <- ret
		}(i)
	}
	// 当第一个人往 channel 中放消息后，接收消息的 receiver 就会从阻塞中被唤醒
	// 一旦 channel 中有消息，则会直接 return 出去
	return <- ch
}

func TestFirstRsponse(t *testing.T) {
	// 输出当前系统中的协程数
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}