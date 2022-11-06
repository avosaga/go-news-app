package main

import (
	"github.com/avosaga/go-news-app/api"
	"github.com/avosaga/go-news-app/config"
	"github.com/avosaga/go-news-app/pkg/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Getting config...")
	conf := config.GetConfig()

	log.Println("Creating API Client...")
	client := &http.Client{Timeout: 10 * time.Second}
	newsApi := api.NewClient(client, conf.ApyKey, 15)

	log.Println("Running server...")
	mux := http.NewServeMux()
	mux.Handle("/assets/", handlers.AssetsHandler)
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/search", handlers.SearchHandler(newsApi))

	log.Printf("Server up and running at https://localhost%s", conf.Port)
	http.ListenAndServe(conf.Port, mux)
}
