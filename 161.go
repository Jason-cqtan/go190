package main

func main() {
	var s []int
	s = append(s, 1)

	var s1 []int
	s1 = make([]int, 0)
	s1 = append(s1, 1)

	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1
}
