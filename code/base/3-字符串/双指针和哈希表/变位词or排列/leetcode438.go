package string_char

// 找到字符串中所有字母异位词

func FindAnagrams(s string, p string) []int {
	result:=make([]int,0)
   if len(p)>len(s){
	   return result
   }
   charArray:=make([]int,26)
   //1 s1，s1 初始化放入到26个字符map中
   for i:=0;i<len(p);i++ {
	   charArray[p[i]-'a']++
	   charArray[s[i]-'a']--
   }
   if allZero(charArray){
	   result=append(result,0)
   }
   for i:=len(p);i<len(s);i++{
	fmt.Println(i)
	fmt.Printf("start:%v ",string((s[i-len(p)])))
	fmt.Printf("end:%v ",string(s[i]))
	   start:=s[i-len(p)]-'a'
	   end:=s[i]-'a'
	   //这个下标的位置 搞清楚
	   charArray[start]++
	   charArray[end]--
	   if allZero(charArray){
		result=append(result,i-len(p)+1)
	   }
   }
   return result
}

func  allZero(count[]int) bool{
   for _,item:=range count{
	   if item!=0{
		   return false
	   }
   }
   return true
}