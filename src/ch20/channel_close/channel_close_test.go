package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 在不确定要放多少数字的时候，如何通知接收者已经放完了？
		// eg：放token（如-1），但是reveiver会很多，而producer不知道其数量
		// 
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 发送数据完毕之后将channel关闭掉
		close(ch)
		// 关闭之后再往 channel 上发数据，会 panic
		// ch <- 11
		wg.Done()
	}()
}


// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 如果通道已经关闭，再从通道中取值
		// 会立即返回通道类型的“零值”
		// 下例多输出一个0
		// for i := 0; i < 11; i++ {
		// 	data := <-ch
		// 	fmt.Println(data)
		// }

		for {
			// ok 返回 false 的时候表示 channel 已经关闭了
			if data, ok := <- ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}