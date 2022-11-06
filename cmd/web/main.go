package main

import (
	"github.com/avosaga/go-news-app/api"
	"github.com/avosaga/go-news-app/config"
	"github.com/avosaga/go-news-app/pkg/handlers"
	"net/http"
	"time"
)

func main() {
	conf := config.GetConfig()

	client := &http.Client{Timeout: 10 * time.Second}
	newsApi := api.NewClient(client, conf.ApyKey, 15)

	mux := http.NewServeMux()
	mux.Handle("/assets/", handlers.AssetsHandler)
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/search", handlers.SearchHandler(newsApi))

	http.ListenAndServe(conf.Port, mux)
}
