package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("length:", len(s))

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slice
	l := s[2:5]
	fmt.Println("slice:", l)
	l = s[:5]
	fmt.Println("slice2", l)
	l = s[2:]
	fmt.Println("slice3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("one line declare:", t)
	twoDimen()
}

func twoDimen() {
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
