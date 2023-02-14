package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	b := base{num: 5}
	fmt.Println(b.describe())
	// container
	co := container{base: base{num: 1}, str: "some name"}
	fmt.Println("num:", co.num)
	fmt.Println("str:", co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("desc:", co.describe())
	fmt.Println("desc:", co.base.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
