package string

func findAnagrams(s string, p string) []int {
    window,need:=make(map[byte]int),make(map[byte]int)
    //window窗口只统计p中有的字符
    left,right,valid:=0,0,0
    res:=make([]int,0)
    for i:=0;i<len(p);i++{
        need[p[i]]++
    }
    for right<len(s){
        ch1:=s[right]
        //错误点：right是在这里自加的
        right++
        if _,ok:=need[ch1];ok{
            window[ch1]++
            if window[ch1]==need[ch1]{
                valid++
            }  
        }
        //如何判断 ceba
        for right-left>=len(p){
            if valid==len(need){
                res=append(res,left)
            }
            ch2:=s[left]
            left++
             if _,ok:=need[ch2];ok{
                 if window[ch2]==need[ch2]{
                     valid--
                 }
                 window[ch2]--
             } 
        }
    }
    return res
}
