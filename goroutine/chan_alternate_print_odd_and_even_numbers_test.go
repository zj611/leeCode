package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//两个协程交替打印奇偶数
func Test_chan_alternate_print_odd_and_even_numbers(t *testing.T) {
	c := make(chan int, 1)

	go func() {
		for i := 1; i <= 100; i++ {
			c <- 1
			if i%2 == 1 {
				fmt.Println("来自协程1", i)
			}
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			<-c
			if i%2 == 0 {
				fmt.Println("来自协程2", i)
			}
		}
	}()

	time.Sleep(3 * time.Second)
}

func TestChanWithTimeout(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("收到信号，监控退出,time=", time.Now().Unix())
				return
			default:
				fmt.Println("goroutine监控中,time=", time.Now().Unix())
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
}

func TestChanWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("收到信号，监控退出,time=", time.Now().Unix())
				return
			default:
				fmt.Println("goroutine监控中,time=", time.Now().Unix())
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(3 * time.Second)
}
