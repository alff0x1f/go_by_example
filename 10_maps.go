package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map after delete:", m)

	fmt.Println("not exist value", m["k2"])

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// one line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("one line map:", n)
}
