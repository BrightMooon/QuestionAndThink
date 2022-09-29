# 切片
## 切片是左闭右开的
```
numbers := []int{0,1,2,3,4,5,6,7,8}  
// 默认下限为 0
fmt.Println("numbers[:3] ==", numbers[:3])
// 结果
numbers[:3] == [0 1 2]
numbers[4:] == [4 5 6 7 8]
```