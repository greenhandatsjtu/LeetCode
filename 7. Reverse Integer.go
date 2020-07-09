package main

import (
	"fmt"
)

func reverse(x int) int {
	var result int32
	base := int32(214748364)
	if x >= 0 && x < 10 || x <= 0 && x > -10 {
		return x
	}
	for x != 0 {
		if result > base || result < -base {
			return 0
		}
		result *= 10
		tmp := x % 10
		result += int32(tmp)
		x /= 10
	}
	return int(result)
}

func main() {
	fmt.Println(reverse(-2147483648))
}
