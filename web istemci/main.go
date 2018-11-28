package main

import (
	"io"
	"net/http"
)

func anaSayfa(cevap http.ResponseWriter, istek *http.Request) {

	cevap.Header().Set("Content Type", "text/html")
	io.WriteString(cevap, `
	<!DOCTYPE html>
	 <html>
		 <body>
		  <h1> Ana Sayfa </h1>
			<a href="/second"> ikinci sayfa için tıkla </a>
			</br>
			<a href="/thirth"> ücüncü sayfa</a>
		</body>	   
	</html>
			 	`)
}

func ikinciSayfa(cevap http.ResponseWriter, istek *http.Request) {

	cevap.Header().Set("Content Type", "text/html")
	io.WriteString(cevap, "ikinci sayfa ")

}
func UcuncuSayfa(cevap http.ResponseWriter, istek *http.Request) {
	cevap.Header().Set("Content Type", "text/html")
	io.WriteString(cevap, " Ucuncu sayfa")
}
func main() {
	http.HandleFunc("/", anaSayfa)
	http.HandleFunc("/second/", ikinciSayfa)
	http.HandleFunc("/thirth/", UcuncuSayfa)
	http.ListenAndServe(":8000", nil)
}
