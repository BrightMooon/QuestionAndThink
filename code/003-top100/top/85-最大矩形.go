package top

func maximalRectangle(matrix [][]byte) int {
	
	return 0
}

// //暴力解法
// /*
// 1. 双循环遍历矩阵
// 2. 定义二维数组表示以当前格子结尾的1的个数
// 3. 第0列结尾的都是1，其他列是前面格子值加1
// 4. 向上扩展行，选择较短的作为宽度
// 5. 更新最大面积
// */
// func maximalRectangle(matrix [][]byte) int {
//     width:=make([][]int,len(matrix))
//     for index:=range width{
//         width[index]=make([]int,len(matrix[0]))
//     }
//     area:=0
//     for row:=0;row<len(matrix);row++{
// 		for col:=0;col<len(matrix[0]);col++{
// 			if matrix[row][col]=='1'{
// 				if col==0{
// 					width[row][col]=1
// 				}else{
// 					width[row][col]=width[row][col-1]+1
// 				}
// 			}else{
// 				width[row][col]=0
// 			}
// 			minWidth:=width[row][col]
// 			for up:=row;up>0;up--{
// 				height:=row-up+1
// 				minWidth=min(minWidth,width[up][col])
// 				area=max(area,minWidth*height)
// 			}
// 		}
// 	}
// 	return area
// }
// func min(a,b int)int{
// 	if a<b {return a}
// 	return b
// }

// func max(a,b int)int{
// 	if a<b {return b}
// 	return a
// }