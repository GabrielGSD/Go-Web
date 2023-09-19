package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Definir a rota "/" e a função que será executada
	// ResponseWriter: escrever a resposta
	// Request: ler a requisição
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá Mundo!")
	})
	// Ler a minha porta 8080 e executar o meu servidor
	fmt.Println(http.ListenAndServe(":8080", nil))
}
