package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	// Создать подразделение
	mux.HandleFunc("POST /departments/")

	// Создать сотрудника в подразделении
	mux.HandleFunc("POST /departments/{id}/employees/")

	// Получить подразделение
	mux.HandleFunc("GET /departments/{id}")

	// Переместить подразделение в другое
	mux.HandleFunc("PATCH /departments/{id}")

	// Удалить подразделение
	mux.HandleFunc("DELETE /departments/{id}")
}
