package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	//"log"
	//"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Endpoint de l'API
	r.HandleFunc("/spots", getSpotsHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Serveur écoutant sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}

func getSpotsHandler(w http.ResponseWriter, r *http.Request) {
	// Connexion à la base de données MySQL à distance
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/yellowaves")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Exécutez la requête SELECT
	rows, err := db.Query("SELECT * FROM yellowavestable")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice pour stocker toutes les destinations
	var allSpots []struct {
		Destination             string `json:"destination"`
		Difficulty              int    `json:"difficulty"`
		DestinationStateCountry string `json:"destinationStateCountry"`
		Photos                  string `json:"photos"`
		SurfBreak               string `json:"surfBreak"`
		PeakSurfSeasonBegins    string `json:"peakSurfSeasonBegins"`
		PeakSurfSeasonEnds      string `json:"peakSurfSeasonEnds"`
	}

	// Parcours des résultats
	for rows.Next() {
		var spot struct {
			Destination             string `json:"destination"`
			Difficulty              int    `json:"difficulty"`
			DestinationStateCountry string `json:"destinationStateCountry"`
			Photos                  string `json:"photos"`
			SurfBreak               string `json:"surfBreak"`
			PeakSurfSeasonBegins    string `json:"peakSurfSeasonBegins"`
			PeakSurfSeasonEnds      string `json:"peakSurfSeasonEnds"`
		}

		err := rows.Scan(&spot.Destination, &spot.DestinationStateCountry, &spot.Difficulty, &spot.SurfBreak, &spot.Photos, &spot.PeakSurfSeasonBegins, &spot.PeakSurfSeasonEnds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Ajout de la destination à la slice
		allSpots = append(allSpots, spot)
	}

	// Convertissez la slice en JSON
	jsonData, err := json.Marshal(allSpots)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Définir le type de contenu de la réponse en JSON
	w.Header().Set("Content-Type", "application/json")

	// Envoyer la réponse JSON
	w.Write(jsonData)
}
