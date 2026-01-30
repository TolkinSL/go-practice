package main

import (
	"dev01/logger"
	"dev01/render"
	"encoding/json"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)
const (
	port = "8080"
)

func main() {
	logger.Init("info", true)

	r := chi.NewRouter()
	r.Post("/api/v1/about_country", Handler)

	log.Info().Msgf("Server start on port %s", port)

	err := http.ListenAndServe(net.JoinHostPort("", port), r)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed start server")
	}
}

type Input struct {
	Countries []string `json:"countries"`
}

type Output struct {
	Results []Result `json:"results"`
}

type Result struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
	Population int `json:"population"`
	Area int `json:"area"`
	Currency string `json:"currency"`
	Languages []string `json:"languages"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("Failed to decode request body")

		return
	}

	if len(input.Countries) == 0 {
		http.Error(w, "Countries field is required", http.StatusBadRequest)
		log.Error().Msg("Countries field is required")

		return
	}

	// Ответ
	var results []Result

	for _, country := range input.Countries {
		result := Result{
			Country: country,
		}

		results = append(results, result)
	} 

	output := Output{
		Results: results,
	}
	render.JSON(w, output, http.StatusOK)
	log.Info().Msgf("Processed counties: %v", input.Countries)
}