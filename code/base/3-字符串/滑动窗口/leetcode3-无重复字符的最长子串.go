package string
func lengthOfLongestSubstring(s string) int {
    left,right:=0,0
    window:=make(map[byte]int)
    res:=0
    for right=0;right<len(s);right++{
        ch1:=s[right]
        window[ch1]++
        //易错点 循环，大于1说明前面出现了重复
        for  window[ch1]>1{
            ch2:=s[left]
            window[ch2]--
            left++
        }
        res=max(res,right-left+1)
    }
    return res
}

func max(a,b int)int{
    if a<b{return b}
    return a
}