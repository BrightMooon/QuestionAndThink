package string_char
//字符串中的变位词==s2包含s1的排列
//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。

/*
结题思路：
1. 26个字母放在一个长数组中充当hash表
2. 
*/
func checkInclusion(s1 string, s2 string) bool {
    if len(s1)>len(s2){
        return false
    }
	charArray:=make([]int,26)
	//1 s1，s1 初始化放入到26个字符map中
	for i:=0;i<len(s1);i++ {
		charArray[s1[i]-'a']++
		charArray[s2[i]-'a']--
        //这里为什么要放出去,起码要等子串全放完，再做判断
		// if allZero(charArray){
		// 	return true
		// }
	}
    if allZero(charArray){
        return true
    }
	for i:=len(s1);i<len(s2);i++{
		charArray[s2[i]-'a']--
		charArray[s2[i-len(s1)]-'a']++
		if allZero(charArray){
			return true
		}
	}
	return false
}
