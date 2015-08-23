package fun2

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

// 简单运算解析器，使用表达式树实现

// 表达式二叉树节点
type binaryNode struct {
	element string // 节点内容
	left    *binaryNode
	right   *binaryNode
}

type token struct {
	flag  string // 标识符类型  integer plus 等
	value int    // 值 操作符使用负数代替
}

type interpreter struct {
	text string // 输入的表达式
	pos  int    // 解析过程中保存的索引
	num  string // 解析过程中保存的数字字符串
}

func (inter *interpreter) getNextToken() token {
	if inter.pos >= len(inter.text) {
		return token{"eof", 0}
	}
	ch := inter.text[inter.pos]
	var t token

	for ch == ' ' {
		inter.pos++
		ch = inter.text[inter.pos]
	}

	isnum := false
	for ch != '+' && ch != '-' && ch != '*' && ch != '/' && ch != ' ' {
		isnum = true
		inter.num += string(ch)
		inter.pos++
		if inter.pos >= len(inter.text) {
			break
		}
		ch = inter.text[inter.pos]
	}

	if isnum {
		v, _ := strconv.Atoi(inter.num)
		t = token{"integer", v}
		inter.num = ""
	} else {
		if ch == '+' {
			t = token{"plus", -1}
		} else if ch == '-' {
			t = token{"cut", -2}
		} else if ch == '*' {
			t = token{"multiply", -3}
		} else if ch == '/' {
			t = token{"divide", -4}
		}

		inter.pos++
	}

	return t
}

func (inter interpreter) eat(flag string, t token) {
	if flag == "op" {
		if t.flag != "plus" && t.flag != "cut" && t.flag != "multiply" && t.flag != "divide" {
			panic("op is not supported")
		}
	} else if t.flag != flag {
		panic("token flag is wrong")
	}
}

func (inter interpreter) expr() int {
	// 栈
	stack := list.New()

	var lastOp string
	for token := inter.getNextToken(); token.flag != "eof"; token = inter.getNextToken() {
		fmt.Printf("token is %s value is %d\n", token.flag, token.value)

		if token.flag == "plus" || token.flag == "cut" || token.flag == "multiply" || token.flag == "divide" {
			// token是操作符
			if lastOp == "" {
				// 第一个操作符
				digitEm := stack.Back()
				stack.Remove(digitEm)
				digit := digitEm.Value.(binaryNode)

				op := binaryNode{token.flag, &digit, nil}
				stack.PushBack(op)
			} else {
				if (lastOp == "multiply" || lastOp == "divide") && (token.flag == "plus" || token.flag == "cut") {
					// 当前的操作符优先级低
					digitEm := stack.Back()
					stack.Remove(digitEm)
					opEm := stack.Back()
					stack.Remove(opEm)

					digit := digitEm.Value.(binaryNode)
					op := opEm.Value.(binaryNode)

					op.right = &digit // 将数字和前一个表达式树绑定

					stack.PushBack(binaryNode{token.flag, &op, nil})
				} else if (token.flag == "multiply" || token.flag == "divide") && (lastOp == "plus" || lastOp == "cut") {
					// 当前的操作符优先级高
					digitEm := stack.Back()
					stack.Remove(digitEm)
					opEm := stack.Back()
					stack.Remove(opEm)

					digit := digitEm.Value.(binaryNode)
					op := opEm.Value.(binaryNode)

					currentNode := binaryNode{token.flag, &digit, nil} // 将数字和当前表达式树绑定
					op.right = &currentNode

					stack.PushBack(op)
				}
			}
			lastOp = token.flag
		} else {
			// 数字入栈
			stack.PushBack(binaryNode{strconv.Itoa(token.value), nil, nil})
		}
	}

	digitEm := stack.Back()
	stack.Remove(digitEm)
	opEm := stack.Back()
	stack.Remove(opEm)
	digit := digitEm.Value.(binaryNode)
	op := opEm.Value.(binaryNode)

	printTree(&op)

	addLastNum(&op, digit.element)

	return calc(op)
}

func printTree(tree *binaryNode) {
	if tree != nil {
		printTree(tree.left)
		fmt.Printf("tree node elemet is %s\n", tree.element)
		printTree(tree.right)
	}
}

func addLastNum(tree *binaryNode, num string) {

	if tree.left != nil || tree.right != nil {
		if tree.left != nil && tree.right != nil {
			addLastNum(tree.left, num)
			addLastNum(tree.right, num)
		} else if tree.left != nil && tree.right == nil {
			tree.right = &binaryNode{num, nil, nil}
		}
	}
}

// 中序遍历表达式树，计算结果
func calc(tree binaryNode) int {

	if tree.left == nil && tree.right == nil {
		v, _ := strconv.Atoi(tree.element)
		return v
	} else {
		left := calc(*tree.left)
		op := tree.element
		right := calc(*tree.right)

		var result int
		if op == "plus" {
			result = left + right
		} else if op == "cut" {
			result = left - right
		} else if op == "multiply" {
			result = left * right
		} else if op == "divide" {
			result = left / right
		}

		return result
	}
}

func SimpleInterpreter() {
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		fmt.Print("calc> ")

		data, _, _ := reader.ReadLine()
		command := string(data)

		if command == "stop" {
			running = false
		} else {
			inter := interpreter{command, 0, ""}
			fmt.Printf("calc result: %d\n", inter.expr())
		}
	}
}
