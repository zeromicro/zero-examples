package main

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/mr"
)

func main() {
	mr.MapReduceVoid(func(source chan<- int) {
		for i := 0; i < 10; i++ {
			source <- i
		}
	}, func(item int, writer mr.Writer[int], cancel func(error)) {
		i := item
		if i == 0 {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(5 * time.Second)
		}
		writer.Write(i)
	}, func(pipe <-chan int, cancel func(error)) {
		for i := range pipe {
			fmt.Println(i)
		}
	})
}
