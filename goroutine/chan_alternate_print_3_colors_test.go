package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//多个协程交替三种颜色

func task(id int, color string, chanA, chanB chan bool, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(chanB)
			return
		default:
			<-chanA
			fmt.Println(fmt.Sprintf("the %d is printing %s", id, color))
			chanB <- true
		}
	}
}

func Test_chan_alternate_print_3_colors(t *testing.T) {
	chanA := make(chan bool)
	chanB := make(chan bool)
	chanC := make(chan bool)

	ctx, cancel := context.WithCancel(context.Background())

	go task(1, "red", chanA, chanB, ctx)
	go task(2, "yellow", chanB, chanC, ctx)
	go task(3, "blue", chanC, chanA, ctx)

	chanA <- true

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

}
