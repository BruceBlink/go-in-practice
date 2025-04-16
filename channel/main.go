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
			
		default:  //channel不够时
			fmt.Println("droped")

		}
		time.Sleep(time.Millisecond * 50)
		
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)  // 模拟处理消息阻塞
	}
}
