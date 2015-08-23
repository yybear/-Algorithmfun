package fun

import (
	"container/list"
	"fmt"
)

// 项目管理图的拓扑排序，使用顶点表示活动网

// 活动顶点
type activityNode struct {
	name         string
	inCount      int   // 活动的前驱节点个数
	days         int   // 活动需要消耗的天数
	adjacent     int   // 相临活动个数
	adjacentNode []int // 相邻活动节点索引
	sTime        int   // 活动最早开始的时间
}

type activityGraph struct {
	count int
	node  []activityNode
}

// 节点在前驱队列里面的结构体
type nodeQueueItem struct {
	index int // 对应到静态邻接表的索引
	sTime int // 活动开始时间，用于排序
}

func enQueue(nodeQueue *list.List, i int, startTime int) {
	// 根据startTime 排序插入
	fmt.Printf("push index %d\n", i)
	insert := 0

	nodeItem := nodeQueueItem{i, startTime}

	em := nodeQueue.Front()
	for em != nil {
		if em.Value.(nodeQueueItem).sTime > startTime {
			nodeQueue.InsertBefore(nodeItem, em)
			insert = 1
			break
		} else {
			em = em.Next()
		}
	}
	if insert == 0 {
		nodeQueue.PushBack(nodeItem)	
	}
}

// 有向图拓扑序列
func topologicalSorting(graph activityGraph, sortedNodes *list.List) bool {
	nodeQueue := list.New()

	for i, node := range graph.node {
		if node.inCount == 0 {
			//找到前驱为0的
			enQueue(nodeQueue, i, node.sTime)
		}
	}
	
	for nodeQueue.Len() > 0 {
		em := nodeQueue.Front()
			
		item := em.Value.(nodeQueueItem)
		
		fmt.Printf("get index %d \n", item.index)
		
		//fmt.Printf("nodeQueue len is %d, ", nodeQueue.Len())
		/*tm :=*/ nodeQueue.Remove(em)
		//fmt.Printf("nodeQueue len is %d remove index %d\n", nodeQueue.Len(), tm.(nodeQueueItem).index)
		
		sortedNodes.PushBack(graph.node[item.index].name)
		// 遍历节点node的所有邻接点，将前驱减一
		for _, nodeIndex := range graph.node[item.index].adjacentNode {
			
			graph.node[nodeIndex].inCount--
			fmt.Printf("node %s, index %d, incount %d\n", graph.node[nodeIndex].name, nodeIndex, graph.node[nodeIndex].inCount)
			if graph.node[nodeIndex].inCount == 0 {
				enQueue(nodeQueue, nodeIndex, graph.node[nodeIndex].sTime)
			}
		}
	}
	
	return graph.count == sortedNodes.Len()
}

func printGraphResult(queue *list.List) {
	fmt.Println("Find result :")
	bs := queue.Front()

	for bs != nil {
		fmt.Printf(",%s", bs.Value.(string))
		bs = bs.Next()
	}
}

func ActivityGraphMain() {
	fmt.Println("活动有向图序列")
	// 静态生成图 使用邻接表方式定义AOV（顶点表示活动网）有向图
	p1 := activityNode{name: "p1", inCount: 0, days: 8, adjacent: 2, adjacentNode: []int{2, 6}, sTime: 0}  // p3 p7
	p2 := activityNode{name: "p2", inCount: 0, days: 5, adjacent: 2, adjacentNode: []int{2, 4}, sTime: 0}  // p3 p5
	p3 := activityNode{name: "p3", inCount: 2, days: 6, adjacent: 1, adjacentNode: []int{3}, sTime: 8}     // p4
	p4 := activityNode{name: "p4", inCount: 1, days: 4, adjacent: 2, adjacentNode: []int{5, 8}, sTime: 14} // p9 p6
	p5 := activityNode{name: "p5", inCount: 1, days: 7, adjacent: 1, adjacentNode: []int{5}, sTime: 5}     // p6
	p6 := activityNode{name: "p6", inCount: 2, days: 7, adjacent: 0, sTime: 18}
	p7 := activityNode{name: "p7", inCount: 1, days: 4, adjacent: 1, adjacentNode: []int{7}, sTime: 8}  // p8
	p8 := activityNode{name: "p8", inCount: 1, days: 3, adjacent: 1, adjacentNode: []int{8}, sTime: 12} // p9
	p9 := activityNode{name: "p9", inCount: 2, days: 4, adjacent: 0, sTime: 18}

	g := activityGraph{9, []activityNode{p1, p2, p3, p4, p5, p6, p7, p8, p9}}
	
	resQueue := list.New()
	res := topologicalSorting(g, resQueue)
	fmt.Printf("有向图无环路：%t\n", res)
	
	printGraphResult(resQueue)
}
