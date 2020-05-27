package main

func subarraysDivByK(A []int, K int) int {
	cache := map[int]int{0: 1}
	var count, sum int
	for _, v := range A {
		sum = ((sum+v)%K + K) % K
		if _, ok := cache[sum]; ok {
			count += cache[sum]
		}
		cache[sum]++
	}
	return count
}
