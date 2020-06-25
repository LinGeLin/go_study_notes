package cancel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	 select {
	 case <- ctx.Done():
		return true
	 default:
		return false
	 }
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if (isCancelled(ctx)) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	// 调用 cancel 方法，取消当前 context 关联的所有节点
	cancel()
	time.Sleep(time.Second * 1)
}