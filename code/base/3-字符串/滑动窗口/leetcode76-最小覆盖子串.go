package string



func minWindow(s string, t string) string {
	//定义needMap 把需要的元素val初始化为1 ,需要的总数为t的长度
	needMap:=make(map[byte]int)
	count:=len(t)
	res :=[]int{0,100000}
	left,right:=0,0
	for i:=0;i<len(t);i++{
		needMap[t[i]]=1
	}
	for right<len(s){
		//1. 从左向右遍历，把字符放入到map中
		if needMap[s[right]]>0{
			count--
		}
		needMap[s[right]]--
		if count==0{
			//2. 收缩：当T中所有的字符都进入了窗口，左指针向右移动，缩小窗口
			for {
				//2.1 碰到需要的字符，停止收缩
				if needMap[s[left]]==0{
					break
				}
				//2.2 碰到不是需要的字符继续收缩
				needMap[s[left]]++
				left++
			}
			if right-left<res[1]-res[0]{
				res=[]int{right,left}
			}
			//3. 移动：窗口移动, 之前的移除了窗口需要累计
			needMap[s[left]]++
			count++
			left++
		}
		right++
	}
	if res[1]>len(s){
		return ""
	}else{
	  return s[res[0]:res[right]]
	}
	
}