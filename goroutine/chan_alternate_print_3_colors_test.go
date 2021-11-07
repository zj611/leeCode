package goroutine

import (
	"fmt"
	"testing"
	"time"
)

//多个协程交替三种颜色
func printRed(ch1, ch2 chan int) {
	<-ch1
	fmt.Println("red")
	ch2 <- 1
}

func printYellow(ch2, ch3 chan int) {
	<-ch2
	fmt.Println("yellow")
	ch3 <- 1
}

func printBlue(ch3, ch1 chan int) {
	<-ch3
	fmt.Println("blue")
	ch1 <- 1
}

func Test_chan_alternate_print_3_colors(t *testing.T) {
	ch1 := make(chan int, 1)
	defer close(ch1)
	ch2 := make(chan int)
	defer close(ch2)
	ch3 := make(chan int)
	defer close(ch3)

	ch1 <- 1
	for i := 0; i < 100; i++ {
		go printRed(ch1, ch2)
		go printYellow(ch2, ch3)
		go printBlue(ch3, ch1)
	}

	time.Sleep(1 * time.Second)
}
