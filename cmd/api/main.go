package main

import (
	"encoding/json"
	"github.com/kumackey/kiriban/kiriban"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syumai/workers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/check/{id}", GetCheck).Methods("GET")

	workers.Serve(r)
}

func GetCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://kumackey.com")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if num >= 10000000 || num <= -10000000 {
		// 10000000以上だとunreachableになる。メモリ使いすぎかもしれん。調査が必要。
		http.Error(w, "number is too large", http.StatusBadRequest)
		return
	}

	resp := GetCheckResponse{
		Number: num,
		Result: kiribanChecker.IsKiriban(num),
		Next:   kiribanChecker.Next(num),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

type GetCheckResponse struct {
	Number int  `json:"number"`
	Result bool `json:"result"`
	Next   int  `json:"next"`
}

var kiribanChecker = MustKiribanChecker()

func MustKiribanChecker() *kiriban.Checker {
	c, err := kiriban.NewChecker(
		kiriban.EnableDigitBasedRoundDetermination(),
	)
	if err != nil {
		panic(err)
	}
	return c
}
