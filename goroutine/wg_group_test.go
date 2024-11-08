package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestWgGroupGoroutine(t *testing.T) {
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup, id int) {
			randT := rand.Intn(100)
			for j := 0; j < randT; j++ {
				time.Sleep(time.Millisecond * 10)
			}
			fmt.Println(fmt.Sprintf("the %d is finished! %d", id, randT))
			wg.Done()
		}(&w, i)
	}

	w.Wait()
	fmt.Println("all goroutine is finished")
}
