package main

import (
	"encoding/json"
	"fmt"

	"log"
	//"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/mux"
)

// func homeLink(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Bienvenue sur le back de Yellowaves !")
// }

// func getRecordByID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	recordID := vars["id"]

// 	for _, record := range records {
// 		if record.ID == recordID {
// 			json.NewEncoder(w).Encode(record)
// 			return
// 		}
// 	}

// 	http.Error(w, "Record not found", http.StatusNotFound)
// }

// func getAllSpotInfo(w http.ResponseWriter, r *http.Request) {

// 	// Convertit les enregistrements en JSON
// 	jsonData, err := json.Marshal(records)
// 	if err != nil {
// 		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Définit le type de contenu de la réponse HTTP comme application/json
// 	w.Header().Set("Content-Type", "application/json")

// 	// Envoie la réponse JSON
// 	w.Write(jsonData)
// }

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homeLink)
// 	router.HandleFunc("/spots", getAllSpotInfo).Methods("GET")
// 	router.HandleFunc("/spots/{id}", getRecordByID).Methods("GET")

//		log.Fatal(http.ListenAndServe(":8080", router))
//	}
func main() {
	// Connexion à la base de données MySQL à distance
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/yellowaves")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Exemple de requête SELECT
	type Spot struct {
		Destination             string `db:"destination"`
		Difficulty              int    `db:"difficulty"`
		DestinationStateCountry string `db:"DestinationStateCountry"`
		Photos                  string `db:"Photos"`
		SurfBreak               string `db:"SurfBreak"`
		PeakSurfSeasonBegins    string `db:"PeakSurfSeasonBegins"`
		PeakSurfSeasonEnds      string `db:"PeakSurfSeasonEnds"`
	}
	// err = db.QueryRow("SELECT * FROM yellowavestable").Scan(&spots.Destination, &spots.DestinationStateCountry, &spots.Difficulty, &spots.SurfBreak, &spots.Photos, &spots.PeakSurfSeasonBegins, &spots.PeakSurfSeasonEnds)
	// if err != nil {
	// 	panic(err)
	// }

	// // Utilisez les données de l'utilisateur ici
	// fmt.Println("Destination:", spots.Destination, "Difficulté:", spots.Difficulty, "DestinationStateCountry:", spots.DestinationStateCountry, "SurfBreak:", spots.SurfBreak, "photos:", spots.Photos, "PeakSurfSeasonBegins:", spots.PeakSurfSeasonBegins, "PeakSurfSeasonEnds:", spots.PeakSurfSeasonEnds)
	rows, err := db.Query("SELECT * FROM yellowavestable")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Créer une slice pour stocker les résultats
	var spotsList []Spot

	// Parcourir les résultats
	for rows.Next() {
		var spots Spot
		err := rows.Scan(
			&spots.Destination,
			&spots.DestinationStateCountry,
			&spots.Difficulty,
			&spots.SurfBreak,
			&spots.Photos,
			&spots.PeakSurfSeasonBegins,
			&spots.PeakSurfSeasonEnds,
		)
		if err != nil {
			log.Fatal(err)
		}
		spotsList = append(spotsList, spots)
	}

	// Gérer les erreurs éventuelles de parcours des résultats
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Encoder en JSON
	jsonData, err := json.MarshalIndent(spotsList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Afficher le JSON
	fmt.Println(string(jsonData))

}
