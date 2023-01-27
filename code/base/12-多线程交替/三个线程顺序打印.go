package main

 
import (
    "fmt"
    "sync"
)

/*
1. 发现明确问题，问题就是问题，不要想其他的
2. chan 阻塞， a协程写入,b协程等待这样，就保证了ab顺序
*/
func Printer(num int) {
    wg := sync.WaitGroup{}
    wg.Add(2)
    // 无buffer 2个g同步发送和接收
    ch := make(chan bool)
    go func() {
        defer wg.Done()
        for i := 1; i <= num; i++ {
            ch <- true
            //奇数
          
                fmt.Println("g1打印:","x")
            
        }
    }()
    go func() {
        defer wg.Done()
        for i := 1; i <= num; i++ {
            <- ch
            //偶数
            
                fmt.Println("g2打印:","y")
            
        }
    }()
    wg.Wait()
}
 
func main() {
    fmt.Println("start ")
    Printer(50)
    fmt.Println("end ")
    

}