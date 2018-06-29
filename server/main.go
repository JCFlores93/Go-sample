package main

import (
	"io"
	"net/http"
)

//* lo usamos como un puntero, para que se guarde en un solo objeto
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola servidor Go!")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
