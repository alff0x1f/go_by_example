package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Сегодня выходные")
	default:
		fmt.Println("Сегодня рабочий день")
	}

	// время
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Время до полудня")
	default:
		fmt.Println("Время после полудня")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Это bool")
		case int:
			fmt.Println("Это int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIAm(true)
	whatIAm(1)
	whatIAm("hey")
}
