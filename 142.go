package main

func main() {
	var m map[int]bool // nil
	_ = m[123]
	var p *[5]string // nil
	for range p {
		_ = len(p)
	}
	var s []int // nil
	_ = s[:]
	//
	s, s[0] = []int{1, 2}, 9// panic s[0] 中s 为nil
}
