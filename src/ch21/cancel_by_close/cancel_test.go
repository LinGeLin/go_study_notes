package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <- cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	// 向 channel 发送一个消息，取消一个任务
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	// 关闭 channel 取消所有任务
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if (isCancelled(cancelCh)) {
					break;
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan)
	// 输出一行

	cancel_2(cancelChan)
	// 输出五行
	time.Sleep(time.Second * 1)
}
