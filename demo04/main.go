package main

import (
	"context"
	"fmt"
	"time"
)

// var wg sync.WaitGroup

// func Prime(num int) bool {
// 	if num == 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= num; i++ {
// 		if num%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func test1(ch chan int) {
// 	for i := 1; i <= 100; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// 	defer wg.Done()
// }

// func test2(ch chan int, ch2 chan int, exitChan chan bool) {
// 	for v := range ch {
// 		if Prime(v) {
// 			ch2 <- v
// 		}
// 	}
// 	exitChan <- true
// 	defer wg.Done()
// }

// func test3(ch2 chan int, exitChan chan bool) {
// 	for v := range ch2 {
// 		fmt.Println(v)
// 	}
// 	exitChan <- true
// 	defer wg.Done()
// }

// func main() {
// 	var ch = make(chan int, 20)
// 	var ch2 = make(chan int, 20)
// 	var exitChan = make(chan bool, 5)
// 	//for i := 0; i < 5; i++ {
// 	wg.Add(1)
// 	go test1(ch)
// 	//}
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go test2(ch, ch2, exitChan)
// 	}
// 	//for i := 0; i < 5; i++ {
// 	wg.Add(1)
// 	go test3(ch2, exitChan)
// 	//}
// 	wg.Add(1)
// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			<-exitChan
// 		}
// 		close(ch2)
// 	}()
// 	wg.Done()

// 	wg.Wait()
// }

func test(msg chan int, ctx context.Context) {
	t := time.Tick(time.Second)
	for range t {
		select {
		case n := <-msg:
			fmt.Println("job", n, "done")
		case <-ctx.Done():
			fmt.Println("return...", ctx.Value(k))
			return
		}
	}
}

type str string

var k str = "name"
var v str = "YueSe"

func main() {
	// ctx, clear := context.WithCancel(context.Background())
	ctx := context.WithValue(context.Background(), k, v)
	ctx, clear := context.WithDeadline(ctx, time.Now().Add(time.Second*5))
	var message = make(chan int)
	// var flag = make(chan bool)
	go test(message, ctx)
	for i := 0; i < 10; i++ {
		message <- i
	}
	clear()
	// flag <- true
	time.Sleep(time.Second)
	fmt.Println("all done")

}
