package main

type Queue []int

func (q *Queue) deque() int {
	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) push(value int) {
	*q = append(*q, value)
}

//BFS inDegree
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses)
	courseMap := make(map[int][]int)
	queue := Queue{}
	for _, prerequisity := range prerequisites {
		courseMap[prerequisity[1]] = append(courseMap[prerequisity[1]], prerequisity[0])
		inDegrees[prerequisity[0]]++
	}
	for i, v := range inDegrees {
		if v == 0 {
			queue.push(i)
		}
	}
	for numCourses != 0 && len(queue) != 0 {
		numCourses--
		course := queue.deque()
		for _, v := range courseMap[course] {
			inDegrees[v]--
			if inDegrees[v] == 0 {
				queue.push(v)
			}
		}
	}
	return numCourses == 0
}

func main() {
	canFinish(2, [][]int{{1, 0}})
}
