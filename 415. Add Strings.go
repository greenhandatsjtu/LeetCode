package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var add bool
	var res string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		temp := a + b
		if add {
			temp += 1
			add = false
		}
		if temp >= 10 {
			add = true
			temp -= 10
		}
		res = strconv.Itoa(temp) + res
	}
	return res
}
