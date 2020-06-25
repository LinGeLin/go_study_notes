package groutine_test

import (
	"testing"
	"fmt"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i< 10; i++ {
		// go 的方法调用都是值传递，传递i的时候都复制了一份
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)

		go func () {
			// i 是共享的，共享就存在竞争条件
			fmt.Println(i)
		}()
		// 输出十行10
	}
	time.Sleep(time.Millisecond * 50)
}