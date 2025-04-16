package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 100)
	for i := range 5 {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		select{
		case c <- rand.Int(): //channel容量足够	
		case <- time.After(time.Millisecond*100):
			fmt.Println("timed out")
		}
	    time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		select {
		case data := <-c:
			fmt.Printf("worker %d got %d\n", w.id, data)
		case<- time.After(time.Microsecond * 10):
			fmt.Println("Break time")
			time.Sleep(time.Second)
		}
	}
}
