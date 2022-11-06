package handlers

import (
	"bytes"
	"github.com/avosaga/go-news-app/api"
	"html/template"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

var tpl, _ = template.ParseFiles("web/templates/index.page.html")
var AssetsHandler = http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets")))

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func SearchHandler(newsApi *api.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parsedUrl, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := parsedUrl.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")

		if page == "" {
			page = "1"
		}

		results, err := newsApi.FetchAll(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &api.Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(newsApi.PageSize))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}
