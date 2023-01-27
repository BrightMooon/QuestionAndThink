package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
    res:=""
    i,j:=len(a)-1,len(b)-1
    carry:=0
    for i>=0||j>=0{
       digitA:= int(a[i]-'0')
       digitB:= int(b[j]-'0')
       i--
       j--
       sum:=digitA+digitB+carry
       if sum>=2{
           sum=sum-2
           carry=1
       }else{
           carry=0
       }
       res=res+strconv.Itoa(sum)
    }
    //循环外处理最后一个进位
    if carry==1{
        res=res+strconv.Itoa(carry)
    }
    return res
}

func main() {
	
	fmt.Print("aaa")

	
}


