package graph

/*
1. 通用思路
找到对应的图，节点，已经和边的关系
广度优先
+ 二维visited数组来保持已经遍历过的节点
+ 定义数组{[-1,0],[1,0],[-1,0][1,0]} 当前坐标和四个方向的坐标相加，得到四个方向移动后的结果
2. 着手点
遍历grid 获取最大值
3. 难点 
对于每个节点，新建一个queue 第一层遍历 queue非空，累加area值，第二层遍历 四个方向的新坐标，且为使用过，放入queue中，同时更新visited数组
4. 易错点

*/
func maxAreaOfIsland(grid [][]int) int {
    row:=len(grid)
    col:=len(grid[0])
    visited:=make([][]bool,row)
    maxArea:=0
    for index:=range visited{
        visited[index]=make([]bool,col)
    }
    for i:=0;i<row;i++{
        for j:=0;j<col;j++{
            if grid[i][j]==1 && !visited[i][j]{
                area:=getArea(grid,visited,i,j)
                maxArea=max(maxArea,area)
            }
        }
    }
    return maxArea
}

func getArea(grid[][]int,visited[][]bool,i,j int)int{
    queue:=make([][]int,0)
    queue=append(queue,[]int{i,j})
    direction:=make([][]int,4)
    visited[i][j]=true
    direction[0]=[]int{-1,0}
    direction[1]=[]int{1,0}
    direction[2]=[]int{0,-1}
    direction[3]=[]int{0,1}
    area:=0
    for len(queue)>0{
        pos:=queue[0]
        queue=queue[1:len(queue)]
        area++
        for _,value:=range direction{
            //当前节点怎么表示
            posX:=pos[0]+value[0]
            posY:=pos[1]+value[1]
            if (posX>=0&&posX<len(grid)&& posY>=0&&posY<len(grid[0])&&grid[posX][posY]==1&&!visited[posX][posY]){
                queue=append(queue,[]int{posX,posY})
                visited[posX][posY]=true
            }
        }

    }
    return area
}

func max(a,b int)int{
    if a<b{return b}
    return a
}