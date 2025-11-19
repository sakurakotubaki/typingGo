package main

import (
	"fmt"
	"typingGo/mylib"
)

func main() {
	p := mylib.Person{Name: "John", Age: 30}
	fmt.Println(p)
	mylib.Say()
}
