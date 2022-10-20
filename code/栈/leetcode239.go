package stackqueue

//滑动窗口的最大值

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

//1 返回队首
func (m *MyQueue) front() int {
	return m.queue[0]
}

//2 返回队尾
func (m *MyQueue) back() int {
	return m.queue[len(m.queue)-1]
}

//3 是否为空
func (m *MyQueue) empty() bool {
	return len(m.queue) == 0
}

//4 新加入元素
/*
4.1 为了维持队列单调从大到小的顺序
4.2  循环比较队尾元素和新加入元素，如果新push的值大于入口元素的数值就讲队列尾部的元素弹出；循环停止条件 新加入元素小于等于队尾元素
*/
func (m *MyQueue) push(val int) {
	//每次push都会做一次循环
	for !m.empty() && m.back() < val {
		//去掉末尾的元素后赋新值
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

//5 弹出元素
//队列非空且将要弹出元素是否为队列出口元素 ,这里不加x==deque.peek()的话，
//当{1,3,-1} ,-3要入队的时候，需要移除1,但在之前因为3入队的时候就把
//1给移除了，此时队头为3，而不是1，x=1,但是队头为3，所以此时不能移除对头，不操作即可
//链接：https://leetcode.cn/problems/sliding-window-maximum/solution/239-hua-dong-chuang-kou-zui-da-zhi-by-we-8hv9/

func (m *MyQueue) pop(val int) {
	if !m.empty() && val == m.front() {
		m.queue = m.queue[1:]
	}
}

//key-1 入队的逻辑 和 滑动窗口一个一个移动的逻辑是不同步的，
//可能入队的时候，窗口的前三个元素都删除了，
//但此时滑动窗口只需要向前移动一个位置
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	queue := MyQueue{}
	//1. 将前K个元素放入到队列中
	for i := 0; i < k; i++ {
		queue.push(nums[i])
	}
	//2. 开始移动窗口，处理入队和出队的情况，并记录结果
	for i := k; i < len(nums); i++ {
		queue.pop(nums[i-k])
		queue.push(nums[i])
		result = append(result, queue.front())
	}
	return result
}
