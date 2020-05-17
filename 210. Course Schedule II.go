package main

//BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	subsequence := make(map[int][]int)
	degree := make([]int, numCourses)
	var queue, result []int
	for _, value := range prerequisites {
		pre := value[1]
		post := value[0]
		degree[post]++
		if subsequence[pre] != nil {
			subsequence[pre] = append(subsequence[pre], post)
		} else {
			subsequence[pre] = []int{post}
		}
	}
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		currentCourse := queue[0]
		queue = queue[1:]
		result = append(result, currentCourse)
		for _, course := range subsequence[currentCourse] {
			degree[course]--
			if degree[course] == 0 {
				queue = append(queue, course)
			}
		}
	}
	if len(result) == numCourses {
		return result
	} else {
		return []int{}
	}
}
