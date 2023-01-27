package main

/*
select 的解决问题:同时响应多个通道的操作

for{
    // 尝试从ch1接收值
    data, ok := <-ch1
    // 尝试从ch2接收值
    data, ok := <-ch2
    …
}
*/

// func main(){
// 	ch:=make(chan int)
// 	ch <- 10 
// 	x := <- ch // 从ch中接收值并赋值给变量x
// 	<-ch  

// }

