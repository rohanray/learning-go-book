package main

import (
	"fmt"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}

var calcMap = map[string]func(int, int) int{
	"+": add,
	"-": minus,
	"*": multiply,
	"/": divide,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		calc, ok := calcMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		out := calc(p1, p2)
		fmt.Println(out)
	}
}
