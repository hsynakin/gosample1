package main

type messagePrinter struct {
	message string
}

func main() {

	mp := messagePrinter{"foo"}
	mp.printMessage()

}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
	println(mp.printMessage)
}
