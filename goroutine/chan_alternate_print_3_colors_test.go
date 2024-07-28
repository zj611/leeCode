package goroutine

import (
	"fmt"
	"testing"
	"time"
)

//多个协程交替三种颜色

func printColor(color string, chA, chB chan int) {
	<-chA
	fmt.Println(color)
	chB <- 1
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
		go printColor("red", ch1, ch2)
		go printColor("yellow", ch2, ch3)
		go printColor("blue", ch3, ch1)
	}

	time.Sleep(1 * time.Second)

	jjj := 27
	fmt.Println(jjj >> 1)
}
