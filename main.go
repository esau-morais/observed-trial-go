package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed static
var static embed.FS

type delivery struct {
	ID     string `json:"id"`
	Parcel string `json:"parcel"`
	Status string `json:"status"`
}

var routes = map[string][]delivery{
	"north": {
		{ID: "N-104", Parcel: "Seed trays", Status: "Out for delivery"},
		{ID: "N-117", Parcel: "Rain barrel", Status: "Delivered"},
	},
	"south": {
		{ID: "S-201", Parcel: "Bench vise", Status: "At depot"},
	},
}

func main() {
	pages, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/deliveries", func(w http.ResponseWriter, r *http.Request) {
		list, ok := routes[r.URL.Query().Get("route")]
		if !ok {
			http.Error(w, "unknown route", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"deliveries": list})
	})
	mux.Handle("GET /", http.FileServerFS(pages))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := "127.0.0.1:" + port
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
