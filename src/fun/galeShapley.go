package fun


// 最为简单的情况，男女个数一样，偏爱列表也是完全的
import (
	"fmt"
)

const COUNT = 3

// 舞伴对象
type partner struct {
	name string
	next int    // 下个选择对象
	current int // -1 表示没有舞伴
	pCount int
	perfect []int
}

// 获取一个自由的男生
func findFreePartner(boys []partner) int {
	for index, boy := range boys {
		if boy.current == -1 {
			return index
		}
	}
	
	return -1
}

func getPerfectPosition(perfect []int, bid int) int {
	for index, id := range perfect {
		if bid == id {
			return index
		}
	}
	
	return 0x7fffffff
}

func isAllMatch(boys []partner) bool {
	for _, boy := range boys {
		if boy.current == -1 {
			return false
		}
	}
	return true
}

func GaleShapley() {
	fmt.Println("boys and girls")
	
	var boys, girls []partner
	boys = make([]partner, 3)
	girls = make([]partner, 3)
	boys[0] = partner{"boy1", 0, -1, COUNT, []int{0, 1, 2}}
	boys[1] = partner{"boy2", 0, -1, COUNT, []int{1, 0, 2}}
	boys[2] = partner{"boy3", 0, -1, COUNT, []int{1, 0, 2}}
	
	girls[0] = partner{"girl1", 0, -1, COUNT, []int{1, 0, 2}}
	girls[1] = partner{"girl2", 0, -1, COUNT, []int{1, 2, 0}}
	girls[2] = partner{"girl3", 0, -1, COUNT, []int{2, 0, 1}}
	
	bid := findFreePartner(boys)
	
	for i:=0; i<9 && bid >= 0; i++ {
		
		boy := &boys[bid]
		
		gid := boy.perfect[boy.next]
		girl := &girls[gid]
		
		if girl.current != -1 {
			// 女生已经有舞伴
			oldBid := girl.current
			
			if getPerfectPosition(girl.perfect, bid) > getPerfectPosition(girl.perfect, oldBid) {
				// 女生喜欢现在的男生
				girl.current = bid
				boy.current = gid
				boys[oldBid].current = -1
			}
		} else {
			boy.current = gid
			girl.current = bid
		}
		boy.next++
		fmt.Printf("boy %d current partner is %d\n", bid, gid)
		
		bid = findFreePartner(boys)
	}
	
	fmt.Printf("isAllMatch: %f", isAllMatch(boys))
}