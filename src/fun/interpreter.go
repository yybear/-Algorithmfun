package fun

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

// 简单运算解析器

type token struct {
	flag string  // 标识符类型  integer plus 等
	value int    // 值 操作符使用负数代替
}

type interpreter struct {
	text string  // 输入的表达式
	pos int      // 解析过程中保存的索引
	num string   // 解析过程中保存的数字字符串
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
		if t.flag != "plus" && t.flag != "cut" && t.flag != "multiply" && t.flag != "divide"  {
			panic("op is not supported")
		}
	} else if t.flag != flag {
		panic("token flag is wrong")
	}
}

func (inter interpreter) expr() int {
	first := inter.getNextToken()
	inter.eat("integer", first)
	
	result := 0
	result = first.value
	
	op := inter.getNextToken()
	
	// 循环支持多数操作
	for op.flag == "plus" || op.flag == "cut" || op.flag == "multiply" || op.flag == "divide" {
		inter.eat("op", op)
		fmt.Printf("op is %s value is %d\n", op.flag, op.value)
		
		second := inter.getNextToken()
		inter.eat("integer", second)
		fmt.Printf("second is %s value is %d\n", second.flag, second.value)
		
		// 做一些操作
		if op.flag == "plus" {
			result += second.value
		} else if op.flag == "cut" {
			result -= second.value
		} else if op.flag == "multiply" {
			result *= second.value
		} else if op.flag == "divide" {
			result /= second.value
		}
		
		op = inter.getNextToken()
	}
	
	return result
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