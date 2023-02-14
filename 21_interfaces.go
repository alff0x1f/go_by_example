package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect_ struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect_) area() float64 {
	return r.width * r.height
}
func (r rect_) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

// Animals
type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "???"
}

type Cow struct {
}

func (c Cow) Speak() string {
	return "Moo!"
}
func main() {
	r := rect_{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
	// animals
	animals := []Animal{Dog{}, Cat{}, Llama{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
