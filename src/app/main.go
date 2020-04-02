package main

import (
	"fmt"
	"net/http"
)

func greeting(mensagemInicial string) string {
	return "<b>" + mensagemInicial + "</b>"
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, greeting("Code.education Rocks!"))
	})

	http.ListenAndServe(":8000", nil)

}
