package main

import (
	"api/src/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API")

	r := route.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
