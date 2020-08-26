package main

import (
	"fmt"
)

type operation interface {
	sum() float64
	sub() float64
	mul() float64
	div() float64
	power() float64
}
type operaends struct {
	x, y float64
}

func (o operaends) sum() float64 {
	return o.x + o.y
}

func (o operaends) sub() float64 {
	return o.x - o.y
}

func (o operaends) mul() float64 {
	return o.x * o.y
}

func (o operaends) div() float64 {
	return o.x / o.y
}

func calculate(o operation) {
	fmt.Println(o.sum())
	fmt.Println(o.sub())
	fmt.Println(o.mul())
	fmt.Println(o.div())
	//fmt.Println(o.power())
}

func main() {
	fmt.Println("ENTER TWO VALUES !")
	num1, num2 := 1, 1
	fmt.Scan(&num1, &num2)
	a := operaends{x: float64(num1), y: float64(num2)}
	calculate(a)
}
