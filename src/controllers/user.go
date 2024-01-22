package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Criar usuário no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Conect()

	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewRepositoryUser(db)
	repository.Create(user)
}

// Listar todos os usuários
func IndexUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listar usuários"))
}

// Buscar usuário
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar o usuário"))
}

// Deletar usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar o usuário"))
}

// Atualizar usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar o usuário"))
}
