// 和尚过河，类似water.go采用穷举法，搜索所有状态
package fun

import (
	"fmt"
	"container/list"
)

var goAction = []itemAction {{2, 0, 0}, {1, 1, 0}, {0, 2, 0}, {1, 0, 0}, {0, 1, 0}}
var backAction = []itemAction {{2, 0, 1}, {1, 1, 1}, {0, 2, 1}, {1, 0, 1}, {0, 1, 1}}

type itemAction struct {
	monks int
	demons int
	direct int // 0 local to remote, 1 remote to local
}

type itemState struct {
	local_monks int
	local_demons int
	remote_monks int
	remote_demons int
	
	currentAction itemAction
}

// 坐船
func (is itemState) move(ia itemAction) (newIs itemState) {
	newIs = itemState{local_monks:is.local_monks, local_demons:is.local_demons, 
		remote_monks:is.remote_monks, remote_demons:is.remote_demons}
	if ia.direct == 0 {
		newIs.local_demons -= ia.demons
		newIs.remote_demons += ia.demons
		newIs.local_monks -= ia.monks
		newIs.remote_monks += ia.monks
	} else if ia.direct == 1 {
		newIs.local_demons += ia.demons
		newIs.remote_demons -= ia.demons
		newIs.local_monks += ia.monks
		newIs.remote_monks -= ia.monks
	}
	
	newIs.currentAction = ia
	return
}

func (is itemState) validate() bool {
	if is.local_demons < 0 || is.local_monks < 0 ||
		is.remote_demons <0 || is.remote_monks < 0 ||
		is.local_demons > 3 || is.local_monks > 3 ||
		is.remote_demons > 3 || is.remote_monks > 3 ||
		(is.local_demons > is.local_monks && is.local_monks != 0 ) || 
		(is.remote_demons > is.remote_monks && is.remote_monks != 0) {
		return false
	}
	return true
}

func (is itemState) isFinal() bool {
	if is.remote_demons == 3 && is.remote_monks ==3 {
		return true
	}
	return false
}

func (is itemState) printState() {
	if is.local_demons == 3 && is.local_monks == 3 {
		fmt.Print("开始：")
	} else if is.currentAction.direct == 0 {
		fmt.Print("过河：")
	} else if is.currentAction.direct == 1 {
		fmt.Print("返回：")
	}
	
	msg := "河边（%d僧尼，%d恶魔），船上（%d僧尼，%d恶魔），对岸（%d僧尼，%d恶魔）\n"
	fmt.Printf(msg, is.local_monks, is.local_demons, is.currentAction.monks, 
		is.currentAction.demons, is.remote_monks, is.remote_demons)
}

// 列表中看当前状态是否已经处理过
func isProcessedRiverState(queue *list.List, current itemState) bool {
	size := queue.Len()
	e := queue.Front()
	for i := 0; i < size; i++ {
		if e != nil {
			bs := e.Value.(itemState)
			if bs.local_demons == current.local_demons && 
				bs.local_monks == current.local_monks && 
				bs.remote_demons == current.remote_demons &&
				bs.remote_monks == current.remote_monks && 
				bs.currentAction.direct == current.currentAction.direct {
				return true
			} else {
				e = e.Next()
			}
		}
	}

	return false
}

func printRiver(queue *list.List) {
	fmt.Println("最终结果 :")
	size := queue.Len()
	bs := queue.Front()
	for i := 0; i < size; i++ {
		if bs != nil {
			bs.Value.(itemState).printState()
			bs = bs.Next()
		}
	}
}

func searchRiverStates(queue *list.List) {
	current := queue.Back().Value.(itemState)
	
	if current.isFinal() {
		// 已经是最后状态了
		printRiver(queue)
	} else {
		if current.currentAction.direct == 0 {
			for i:=0; i<5; i++ {
				next := current.move(backAction[i])
				//next.printState()
				if next.validate() && !isProcessedRiverState(queue, next) {
					queue.PushBack(next)
					searchRiverStates(queue)
					queue.Remove(queue.Back())
				}
			}	
		} else {
			for i:=0; i<5; i++ {
				next := current.move(goAction[i])
				//next.printState()
				if next.validate() && !isProcessedRiverState(queue, next) {
					queue.PushBack(next)
					searchRiverStates(queue)
					queue.Remove(queue.Back())
				}
			}		
		}
	}
}

func DemonsAndMonks() {
	// 初始状态
	initState := itemState{3, 3, 0, 0, itemAction{0, 0, 1}}
	
	queue := list.New()
	queue.PushBack(initState)
	
	searchRiverStates(queue)
}