package main

import (
	"fmt"
)

// func main() {
// 	var sum int
// 	var wg sync.WaitGroup
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			sum++
// 			wg.Done()
// 		}()
// 		wg.Wait()
// 	}
// 	fmt.Println(sum)
// }

func main() {
	var sum int
	var blocking (chan bool)
	for i := 0; i < 100; i++ {
		go func() {
			sum++
			blocking <- true
		}()
		<-blocking
	}
	fmt.Println(sum)
}
