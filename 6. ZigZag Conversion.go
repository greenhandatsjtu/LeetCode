package main

import "strings"

//找规律
func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) {
		return s
	}
	var tmp []string
	var index, stride int
	chars := strings.Split(s, "")
	length := len(chars)
	for i := 0; i < numRows; i++ {
		if i >= length {
			break
		}
		index = i
		tmp = append(tmp, chars[index])
		stride = 2*(numRows-index%(2*numRows-2)) - 2
		for {
			if stride != 0 {
				index += stride
				stride = 2*numRows - 2 - stride
			} else {
				stride = 2*numRows - 2 - stride
				continue
			}
			if index >= length {
				break
			}
			tmp = append(tmp, chars[index])
		}
	}
	return strings.Join(tmp, "")
}

func main() {
	println(convert("A", 1))
}
