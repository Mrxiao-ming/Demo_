package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
)

//var sum int32
//
//type testNice struct {
//	ii int32
//}
//
//func myFunc(i interface{}) {
//	n := i.(*testNice).ii
//	atomic.AddInt32(&sum, n)
//	fmt.Printf("run with %d\n", n)
//}
//
//func demoFunc() {
//	time.Sleep(10 * time.Millisecond)
//	fmt.Println("Hello World!")
//}
//
//func main() {
//	defer ants.Release()
//
//	runTimes := 1000
//
//	// Use the common pool.
//	var wg sync.WaitGroup
//	syncCalculateSum := func() {
//		demoFunc()
//		wg.Done()
//	}
//	for i := 0; i < runTimes; i++ {
//		wg.Add(1)
//		_ = ants.Submit(syncCalculateSum)
//	}
//	wg.Wait()
//	fmt.Printf("running goroutines: %d\n", ants.Running())
//	fmt.Printf("finish all tasks.\n")
//
//	// Use the pool with a function,
//	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
//	p, _ := ants.NewPoolWithFunc(3, func(i interface{}) {
//		myFunc(i)
//		wg.Done()
//	})
//
//	defer p.Release()
//	// Submit tasks one by one.
//	for i := 0; i < runTimes; i++ {
//		wg.Add(1)
//		_ = p.Invoke(&testNice{
//			ii: int32(i),
//		})
//	}
//	wg.Wait()
//	fmt.Printf("running goroutines: %d\n", p.Running())
//	fmt.Printf("finish all tasks, result is %d\n", sum)
//}

func main() {
	var (
		wg  sync.WaitGroup
		sum int32
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		pool, _ := ants.NewPoolWithFunc(3, func(i interface{}) {
			atomic.AddInt32(&sum, 1)
			fmt.Println(i)
		})
		defer func() {
			for {
				if pool.Running() == 0 && pool.Waiting() == 0 {
					pool.Release()
					return
				} else {
					continue
				}

			}
		}()

		for i := 1; i <= 50; i++ {
			_ = pool.Invoke(i)
		}
	}()
	wg.Wait()
	fmt.Println("goroutine num:", sum)
}

func TestAA() {
	var (
		a = make(chan int, 10)
		b = make(chan int, 5)
		//c = make(chan int)
	)
	go func() {
		//defer close(a)
		//defer close(b)
		for index, val := range []int{1, 2, 3, 4, 5} {
			a <- val
			b <- index
		}
	}()

	for {
		select {
		case av, ok := <-a:
			if !ok {
				continue
				a = nil
			}
			fmt.Println(av)
		case bv, ok := <-b:
			if !ok {
				continue
				b = nil
			}
			fmt.Println(bv)
			//case _, ok := <-c:
			//	if ok {
			//		fmt.Println("return")
			//		return
			//	}
			//}
		}
	}
}

func ExecCommandWithResult(command string) string {
	out, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil && !strings.Contains(err.Error(), "exit status") {
		log.Println("err: " + err.Error())
		return ""
	}
	return string(out)
}

//func main() {
//	//var (
//	//	port = ":8080"
//	//	//env = flag.String("env", "dev", "请输入运行环境:\n  fat:测试环境\n uat:预上线环境\n product:正式环境 \n")
//	//)
//	//
//	//// mysqlConnect := connections.NewMysqlConnection()
//	//
//	//flag.Parse()
//	//
//	//fmt.Println("start time:", time.Now().Format(string(enum.FormatTime)))
//	//gin.SetMode(gin.DebugMode)
//	//gin.DisableConsoleColor()
//	//engine := gin.New()
//	//engine.Use(route.Cors())
//	//route.InitRoute(engine)
//	//fmt.Println("end time:", time.Now().Format(string(enum.FormatTime)))
//	//_ = engine.Run(port)
//
//}
