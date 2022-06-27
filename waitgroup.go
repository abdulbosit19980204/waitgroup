package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	//goroutine larni anonim funkisya bilan chaqirish

	go func() {
		display("Hello")
		wg.Done()
	}()

	go func() {
		display("How are you")
		wg.Done()
	}()

	//	go display("Hello", &wg)
	//go display("How are you", &wg)
	//time.Sleep(2 * time.Second)
	//fmt.Scanln()
	wg.Wait()
	fmt.Println("completed...")
}
func display(input string) {
	//func display(input string, wg *sync.WaitGroup) {
	for i := 1; i <= 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
	//wg.Done()

}
