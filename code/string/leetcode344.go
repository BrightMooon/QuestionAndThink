package string
/**
反转字符串
*/

//左右双指针替换

func reverseString(s []byte)  {
	left:=0
    right:=len(s)-1
    for left<right{
        s[left],s[right]=s[right],s[left]
        left++
        right--
    }
}