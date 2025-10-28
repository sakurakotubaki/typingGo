package main

import "fmt"

type Human interface {
	Say()
}

type Person struct {
	Name string
}

// My Code
// interface ロジックがない
type Tool interface {
	Plus()
}

// 構造体を定義
type Driver struct {
	Brand string
}

// 引数に構造体、右にinterface名
func (d Driver) Plus() {
	fmt.Println(d.Brand)
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

func main() {
	var mike Human = Person{"Mike"}
	mike.Say()
	var car Tool = Driver{"Toyota"}
	car.Plus()
}
