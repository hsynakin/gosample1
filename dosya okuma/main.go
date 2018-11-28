package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	dosyaBilgi, err := ioutil.ReadFile("testt.txt")
	if err == nil {
		fmt.Println("dosyadan okunun veri :", string(dosyaBilgi))

	}
	if err != nil {
		fmt.Println("Hata mesajı : ", err)
	}
	dosya, _ := os.Create("test2.txt")
	dosya.WriteString(`go go go`)
}
