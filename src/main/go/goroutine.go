package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	a := make(chan bool)
	go hello(a)
	//time.Sleep(1 * time.Second)
	<-a
	fmt.Println("main function")
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v, "from ch")
		time.Sleep(200 * time.Millisecond)

	}

	ch3 := make(chan string, 3)
	ch3 <- "naveen"
	ch3 <- "paul"
	fmt.Println("capacity is", cap(ch3))
	fmt.Println("length is", len(ch3))
	fmt.Println("read value", <-ch3)
	fmt.Println("new length is", len(ch3))

	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")

	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

}

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
