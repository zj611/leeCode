package goroutine

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*
  建立3个协程任务，当任意1个协程完成时，通知其它2个协程立马结束退出
*/

func workerTask(id int, done chan bool, ctx context.Context) {
	// 随机工作0～10s
	randT := rand.Intn(10)
	//timer := time.NewTimer(time.Second * time.Duration(randT))
	fmt.Println(fmt.Sprintf("the %v is working， %v", id, randT))
	for i := 0; i < randT; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("the %v is ending", id))
			return
		default:
			time.Sleep(time.Second)
		}
	}
	done <- true
	fmt.Println(fmt.Sprintf("the %v is ending", id))
}

func TestChanContextClose(t *testing.T) {
	ctx, concel := context.WithCancel(context.Background())
	done := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(id int, ch chan bool) {
			workerTask(id, ch, ctx)
		}(i, done)
	}

	<-done
	concel()
	time.Sleep(10 * time.Second)
	close(done)
}

func workerTask1(id int, done chan bool, ctx context.Context) {
	t := rand.Intn(10)
	fmt.Println(fmt.Sprintf("the %d started! %d", id, t))
	for i := 0; i < t; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("the %d stopped!", id))
			return
		default:
			time.Sleep(time.Second)
		}
	}
	done <- true
	fmt.Println(fmt.Sprintf("the %d stopped!", id))
}

func TestChanContextClose1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)
	for i := 0; i < 3; i++ {
		go workerTask1(i, done, ctx)
	}
	<-done
	cancel()

	time.Sleep(10 * time.Second)
	close(done)

}
