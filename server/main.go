package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type OPARequest struct {
	Input Input `json:"input"`
}

type Input struct {
	User   string   `json:"user"`
	Path   []string `json:"path"`
	Method string   `json:"method"`
}

type OPAResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	Allow bool `json:"allow"`
}

func CheckPolicy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathList := strings.Split(r.URL.Path, "/")
		user, _, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}
		input := Input{
			User:   user,
			Path:   pathList[1:],
			Method: r.Method,
		}
		opaReq := OPARequest{Input: input}
		jsonInput, err := json.Marshal(opaReq)
		if err != nil {
			http.Error(w, "can't marshal json", http.StatusInternalServerError)
			return
		}
		body := bytes.NewReader(jsonInput)
		res, err := http.Post("http://127.0.0.1:8181/v1/data/authz", "application/json", body)
		if err != nil {
			http.Error(w, "can't connect to opa server", http.StatusInternalServerError)
			return
		}
		opaResponse := OPAResponse{}
		err = json.NewDecoder(res.Body).Decode(&opaResponse)
		if err != nil {
			http.Error(w, "can't decode json", http.StatusInternalServerError)
			return
		}
		if opaResponse.Result.Allow {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	})

}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.BasicAuth("Authentication Failed", map[string]string{
		"alice": "demo",
		"bob":   "demo",
		"david": "demo",
	}))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	r.Route("/salary/{username}", func(r chi.Router) {
		r.Use(CheckPolicy)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("300K US$"))
		})
	})

	log.Println("Server is running on port 8081")
	log.Fatalln(http.ListenAndServe(":8081", r))
}
