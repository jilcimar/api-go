package main

import (
	"api/src/config"
	"api/src/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	fmt.Println(config.StringConexction)
	fmt.Println("Rodando API")

	r := route.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
