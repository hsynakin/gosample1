package main

func main() {
	result, tt := add(1, 3, 5, 6, 7)
	println(tt, result)
}
func add(terms ...int) (int, string) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result, "Sonuç  = "

}
