package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 четное")
	} else {
		fmt.Println("7 нечетное")
	}

	if 8%4 == 0 {
		fmt.Println("8 делится на 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "негативное")
	} else if num < 10 {
		fmt.Println(num, "из одного числа")
	} else {
		fmt.Println(num, "из нескольких цифр")
	}

}
