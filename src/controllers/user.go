package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Criar usuário no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		responses.Error(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Error(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conect()

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	repository := repositories.NewRepositoryUser(db)
	userID, erro := repository.Create(user)

	if erro != nil {
		responses.Error(w, http.StatusInternalServerError, erro)
		return
	}

	user.ID = userID

	responses.JSON(w, http.StatusCreated, user)
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
