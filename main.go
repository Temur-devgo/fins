package main

import (
	"encoding/json"
	"fins-api-services/db"
	"fins-api-services/services"
	"fins-api-services/structures"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	db.InitDB(dbDriver, dbUser+":"+dbPass+"@"+dbHost+"/"+dbName)

	router := mux.NewRouter()
	router.HandleFunc("/history/{userId}", getHistory).Methods("GET")
	router.HandleFunc("/expense/", expense).Methods("POST")
	router.HandleFunc("/income/", income).Methods("POST")

	log.Println("Connection success!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services.GetHistory(userId))
}

func expense(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var expenseForm structures.PostForm
	err := decoder.Decode(&expenseForm)
	if err != nil {
		panic(err)
	}

	response := services.Expense(expenseForm.UserId, expenseForm.ToUserId, expenseForm.Sum)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func income(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var incomeForm structures.PostForm
	err := decoder.Decode(&incomeForm)
	if err != nil {
		panic(err)
	}

	response := services.Income(incomeForm.UserId, incomeForm.ToUserId, incomeForm.Sum)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
