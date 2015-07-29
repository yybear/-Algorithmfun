package fun

import (
	"fmt"
	"strconv"
	"strings"
)

// google 方程式

type tagCharItem struct {
	c       byte
	value   int
	leading bool
}

type tagCharValue struct {
	value int
	used  bool
}

func searchEq(chars []tagCharItem, values []tagCharValue, index int) {

	if index > 8 {
		data1s := []string{strconv.Itoa(chars[0].value), strconv.Itoa(chars[0].value),
			strconv.Itoa(chars[0].value), strconv.Itoa(chars[1].value),
			strconv.Itoa(chars[2].value), strconv.Itoa(chars[3].value)}
		s1 := strings.Join(data1s, "")

		data2s := []string{strconv.Itoa(chars[4].value), strconv.Itoa(chars[2].value),
			strconv.Itoa(chars[2].value), strconv.Itoa(chars[4].value),
			strconv.Itoa(chars[5].value), strconv.Itoa(chars[6].value)}
		s2 := strings.Join(data2s, "")

		data3s := []string{strconv.Itoa(chars[1].value), strconv.Itoa(chars[2].value),
			strconv.Itoa(chars[3].value), strconv.Itoa(chars[7].value),
			strconv.Itoa(chars[2].value), strconv.Itoa(chars[8].value)}
		s3 := strings.Join(data3s, "")

		d1, _ := strconv.Atoi(s1)
		d2, _ := strconv.Atoi(s2)
		d3, _ := strconv.Atoi(s3)

		if d1-d2 == d3 {
			fmt.Println(s1 + "-" + s2 + "=" + s3)
		}
		return
	}

	for i := 0; i < 10; i++ {
		if (values[i].value == 0 && chars[index].leading) || values[i].used {
			continue
		}

		chars[index].value = values[i].value
		values[i].used = true

		searchEq(chars, values, index+1)

		values[i].used = false
	}
}

func GoogleEq() {
	fmt.Println("Google Equation")

	//"wwwdot-google=dotcom"

	var chars = []tagCharItem{{'w', -1, true}, {'d', -1, true}, {'o', -1, false},
		{'t', -1, false}, {'g', -1, true}, {'l', -1, false}, {'e', -1, false},
		{'c', -1, false}, {'m', -1, false}}

	var values = []tagCharValue{{0, false}, {1, false}, {2, false}, {3, false},
		{4, false}, {5, false}, {6, false}, {7, false}, {8, false}, {9, false}}

	searchEq(chars, values, 0)
}
