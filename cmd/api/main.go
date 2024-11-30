package main

import (
	"encoding/json"
	"github.com/kumackey/kiriban/internal/oapi"
	"github.com/syumai/workers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ping", http.HandlerFunc(GetPing))

	workers.Serve(mux)
}

func GetPing(w http.ResponseWriter, r *http.Request) {
	resp := oapi.Pong{
		Ping: "pong!",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
