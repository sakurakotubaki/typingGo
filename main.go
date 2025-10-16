package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

// *がないとスコープの中でした変更は反映されない
// func changeVertex(v Vertex) {
// 	v.X = 1000
// 	fmt.Println("scop in value", v)
// }

// *があるとポインターを渡すことができる
// func changeVertex(v *Vertex) {
// 	v.X = 1000
// 	fmt.Println("scop in value", v)
// }

// (*v).X = 1000 と同じ意味
func changeVertex(v *Vertex) {
	(*v).X = 1000
	fmt.Println("(*v)で自動でポインタの実態を指す", v)
}

func main() {
	// v := Vertex{X: 1, Y: 2, S: "hello"}
	// fmt.Println(v)
	// v.X = 10
	// fmt.Println(v.X, v.Y)

	// v6 := new(Vertex)
	// fmt.Println(v6)サイレントTV

	// newで作成するよりも「&Vertex{}」のようにアドレスをつけた状態で宣言した方がポインタが返ってくることが明示的なためよく使われる。
	// v7 := &Vertex{}
	// fmt.Printf("好きだった%s より %d つけvマックス\n", v7.S, v7.X)

	// v := Vertex{X: 1, Y: 2, S: "hello"}
	v := &Vertex{X: 1, Y: 2, S: "ポインターで値を渡したよ☺️ "}
	changeVertex(v)
	fmt.Println(v)
}
