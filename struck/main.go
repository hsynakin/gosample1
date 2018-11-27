package main

type myStruct struct {
	myField string
}

func main() {
	foo := myStruct{"bar"}
	println(foo.myField)

}
