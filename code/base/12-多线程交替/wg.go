package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x int64
// var wg sync.WaitGroup

// func add(){
// 	for i := 0; i <50000; i++ {
// 		x=x+1
// 	}
// 	wg.Done()
// }

// func  main(){
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Print(x)
// }