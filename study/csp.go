package study

import (
	"fmt"
	"sync"
)

/*
	实现两个goroutine交替执行打印数字和字母,效果如下:
		12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ…
*/
func AlternatingPrinting1() {
	var (
		cha1   = make(chan int, 1)
		cha2   = make(chan int)
		finish = make(chan int, 1)
	)
	cha1 <- 1
	finish <- 1
	go func() {
		for i := 1; i < 26; i += 2 {
			<-cha1
			fmt.Print(i)
			fmt.Print(i + 1)
			cha2 <- 1
		}
	}()
	go func() {
		for i := 'A'; i <= 'Z'; i += 2 {
			<-cha2
			fmt.Print(string(i))
			fmt.Print(string(i + 1))
			cha1 <- 1
		}
		finish <- 1
	}()
	<-finish
}

func AlternatingPrinting2() {
	var (
		flag = make(chan int, 1)
		wg   = sync.WaitGroup{}
	)
	go func() {
		wg.Add(1)
		for i := 1; i < 26; i += 2 {
			flag <- 1
			fmt.Print(i)
			fmt.Print(i + 1)
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 'A'; i <= 'Z'; i += 2 {
			<-flag
			fmt.Print(i)
			fmt.Print(i + 1)
			flag <- 1
		}
		wg.Done()
	}()
	wg.Wait()
}
