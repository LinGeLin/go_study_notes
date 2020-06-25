package share_mem

import (
	"testing"
	"time"
	"sync"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		// 没有并发的保护，最终输出counter将小于5000
		go func() {
			counter ++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		// 没有并发的保护，最终输出counter将小于5000
		go func() {
			defer func () {
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mt sync.Mutex
	counter := 0;
	for i:=0; i<5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mt.Unlock()
			}()
			mt.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}