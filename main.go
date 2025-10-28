package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Movie struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Overview  string `json:"overview"`
	PosterURL string `json:"posterurl"`
	Genres    string `json:"genres"`
}

func main() {
	_ = godotenv.Load(".env")
	apiURL := os.Getenv("API_URL")
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8113"
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(apiURL + "/movies")
		if err != nil {
			http.Error(w, "Erro ao buscar filmes", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var movies []Movie
		if err := json.NewDecoder(resp.Body).Decode(&movies); err != nil {
			http.Error(w, "Erro ao decodificar resposta", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, movies)
	})

	log.Printf("Servidor rodando em http://localhost:%s", appPort)
	http.ListenAndServe(":"+appPort, nil)
}
