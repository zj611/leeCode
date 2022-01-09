package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestCloseGoroutine(t *testing.T) {
	//ctx,cancle := context.WithCancel(context.Background())
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx,cancle := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	ch := make(chan int, 0)
	defer cancle()

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- i:
				time.Sleep(time.Millisecond * 100)
				fmt.Println("ch <- ", i)
				//default:

			}
		}
	}(ctx)

	//for {
	//	if v,ok := <-ch; ok{
	//		fmt.Println("rev ", v)
	//	}else {
	//		break
	//	}
	//}

	for v := range ch {
		fmt.Println("rev ", v)
		//if v == 10 {
		//	cancle()
		//	break
		//}
	}

	time.Sleep(5 * time.Millisecond)
}
