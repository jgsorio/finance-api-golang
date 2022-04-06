package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func main() {
	http.HandleFunc("/", getTransactions)
	http.HandleFunc("/create", createATransaction)
	http.ListenAndServe(":8000", nil)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var transactions = Transactions{
		Transaction{
			Title:     "Salario",
			Amount:    1257.00,
			Type:      1,
			CreatedAt: time.Now(),
		},
	}

	json.NewEncoder(w).Encode(transactions)
}

func createATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	res := Transactions{}
	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &res)
	fmt.Println(res)
}
