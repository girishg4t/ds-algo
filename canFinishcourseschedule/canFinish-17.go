package canfinishcourseschedule

import "fmt"

func CanFinish(numCourses int, prerequisites [][]int) bool {
	courcesmap := make(map[int][]int)

	for _, n := range prerequisites {
		courcesmap[n[0]] = append(courcesmap[n[0]], n[1])
	}
	visitset := make(map[int]bool)
	fmt.Println(courcesmap)
	for _, crs := range courcesmap {
		if !dfs(courcesmap, visitset, crs[0]) {
			return false
		}
	}
	return true
}

func dfs(cources map[int][]int, visitset map[int]bool, start int) bool {
	val, ok := cources[start]
	if !ok {
		return true
	}
	if visitset[start] {
		return false
	}
	visitset[start] = true
	for _, cource := range val {
		if !dfs(cources, visitset, cource) {
			return false
		}
	}
	delete(visitset, start)
	delete(cources, start)
	return true
}
