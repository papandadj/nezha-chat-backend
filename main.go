package main

func main() {
	a := A{A: "a", B: &ErrorB{A: "A", B: "B"}}
	test(a)
}

type A struct {
	A string
	B *ErrorB
}

type ErrorB struct {
	A string
	B string
}

func test(a interface{}) {

}
