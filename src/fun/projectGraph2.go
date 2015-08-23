package fun

import (
	"container/list"
	"fmt"
)

// 项目管理图的搜寻关键线路，使用边表示活动网，事先需要进行拓扑排序

// 事件顶点
type eventNode struct {
	inCount int            // 活动的前驱节点个数
	sTime   int            // 事件最早开始时间
	eTime   int            // 事件最晚开始时间
	edge    []activityEdge // 相邻边
}

// 活动边
type activityEdge struct {
	nodeIndex int // 活动边终点顶点索引
	name      string
	duty      int // 活动边持续时间（权重）
}

type graph struct {
	count int
	node  []eventNode
	
}

// 计算顶点的最早开始时间
func calcESTime(g *graph, sortedNode *list.List) {
	// 初始化第一个顶点的最早开始时间为0
	g.node[0].sTime = 0
	
	/*size := sortedNode.Len()
	em := sortedNode.Front()
	
	for i := 0; i < size; i++ {
		if em != nil {
			for 
			
			em = em.Next()
		}
	}*/
	
	fmt.Println("test")
}
