package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range over map
	dict := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range dict {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// range over map, only keys
	for k := range dict {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
