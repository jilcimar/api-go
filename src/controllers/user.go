package controllers

import "net/http"

// Criar usuário no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando o usuário"))
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
