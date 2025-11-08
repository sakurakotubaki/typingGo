package main

import "fmt"

// update1111
// update2222
// update3333

type Person struct {
	Name string
}

type Hoge struct {
	Name string
}

func (p *Person) Say() {
	p.Name = "hoge"
	// %pを使うとポインタのアドレスが表示される
	fmt.Printf("%p\n", p)
}

func main() {
	p := Person{Name: "fuga"}
	p.Say()
	fmt.Println(p.Name)
}
