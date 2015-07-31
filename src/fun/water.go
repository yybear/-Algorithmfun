package fun

import (
	"container/list"
	"fmt"
)

// 三个水桶等分8升水

type action struct {
	from  int
	to    int
	water int
}

var bucket = []int{8, 5, 3}

// 水桶状态 定义
type bucketState struct {
	buckets   []int
	curAction action
}

// 判断是否可以倒水
func (b bucketState) canDump(from, to int) bool {
	// 不能是同一个桶，from桶不能为空，to桶不能满
	if from != to && b.buckets[from] != 0 && b.buckets[to] < bucket[to] {
		return true
	}

	return false
}

func (b bucketState) isFinal() bool {
	if b.buckets[0] == 4 && b.buckets[1] == 4 {
		return true
	}

	return false
}

func (b bucketState) printStates() {
	fmt.Printf("dump %d water from %d to %d, buckets water states is: %d %d %d \n",
		b.curAction.water, b.curAction.from, b.curAction.to, b.buckets[0], b.buckets[1], b.buckets[2])
}

// 倒水
func (b bucketState) dumpWater(from, to int, bs *bucketState) {
	left := bucket[to] - bs.buckets[to]
	var water int
	if bs.buckets[from] <= left {
		// from的水可以全部倒入to
		water = bs.buckets[from]
		bs.buckets[to] = bs.buckets[to] + water
		bs.buckets[from] = 0
	} else {
		// 只能倒一部分
		water = left
		bs.buckets[to] = bucket[to]
		bs.buckets[from] = bs.buckets[from] - water
	}

	bs.curAction = action{from, to, water}
}

func printResult(queue *list.List) {
	fmt.Println("Find result :")
	size := queue.Len()
	bs := queue.Front()
	for i := 0; i < size; i++ {
		if bs != nil {
			bs.Value.(bucketState).printStates()
			bs = bs.Next()
		}
	}
}

// 列表中看当前状态是否已经处理过
func isProcessedState(queue *list.List, current bucketState) bool {
	size := queue.Len()
	e := queue.Front()
	for i := 0; i < size; i++ {
		if e != nil {
			bs := e.Value.(bucketState)
			if bs.buckets[0] == current.buckets[0] && bs.buckets[1] == current.buckets[1] && bs.buckets[2] == current.buckets[2] {
				return true
			} else {
				e = e.Next()
			}
		}
	}

	return false
}

func searchStateOnAction(queue *list.List, current bucketState, from, to int) {
	canDump := current.canDump(from, to)
	if canDump {
		next := bucketState{buckets: []int{current.buckets[0], current.buckets[1], current.buckets[2]}}
		current.dumpWater(from, to, &next)
		if !isProcessedState(queue, next) {
			queue.PushBack(next)
			searchStates(queue)
			queue.Remove(queue.Back())
		}
	}
}

func searchStates(queue *list.List) {
	current := queue.Back().Value.(bucketState)
	if current.isFinal() {
		// 已经是最后状态了
		printResult(queue)
	} else {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				searchStateOnAction(queue, current, j, i)
			}
		}
	}
}

func DumpWater() {
	fmt.Println("water dump")
	queue := list.New()

	bstateInit := bucketState{[]int{8, 0, 0}, action{-1, 0, 8}} // 初始化
	queue.PushBack(bstateInit)

	searchStates(queue)
}
