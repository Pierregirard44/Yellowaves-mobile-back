package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Structure pour représenter un enregistrement
type SurfRecord struct {
	ID          string       `json:"id"`
	Fields      RecordFields `json:"fields"`
	CreatedTime string       `json:"createdTime"`
}

// Structure pour représenter les champs d'un enregistrement
type RecordFields struct {
	SurfBreak               []string `json:"Surf Break"`
	DifficultyLevel         int      `json:"Difficulty Level"`
	Destination             string   `json:"Destination"`
	Geocode                 string   `json:"Geocode"`
	Influencers             []string `json:"Influencers"`
	MagicSeaweedLink        string   `json:"Magic Seaweed Link"`
	Photos                  []Photo  `json:"Photos"`
	PeakSurfSeasonBegins    string   `json:"Peak Surf Season Begins"`
	DestinationStateCountry string   `json:"Destination State/Country"`
	PeakSurfSeasonEnds      string   `json:"Peak Surf Season Ends"`
	Address                 string   `json:"Address"`
}

// Structure pour représenter les photos
type Photo struct {
	ID         string     `json:"id"`
	URL        string     `json:"url"`
	Filename   string     `json:"filename"`
	Size       int        `json:"size"`
	Type       string     `json:"type"`
	Thumbnails Thumbnails `json:"thumbnails"`
}

// Structure pour représenter les miniatures des photos
type Thumbnails struct {
	Small Small `json:"small"`
	Large Large `json:"large"`
	Full  Full  `json:"full"`
}

// Structure pour représenter une miniature
type Small struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Structure pour représenter une grande image
type Large struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Structure pour représenter une image en taille réelle
type Full struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue sur le back de Yellowaves !")
}

func getSpotList(w http.ResponseWriter, r *http.Request) {
	// Simule la récupération des enregistrements depuis votre source de données (base de données, fichier, etc.)
	records := []SurfRecord{
		{
			Fields: RecordFields{
				SurfBreak:       []string{"Reef Break"},
				DifficultyLevel: 4,
			},
		},
	}

	// Convertit les enregistrements en JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	// Définit le type de contenu de la réponse HTTP comme application/json
	w.Header().Set("Content-Type", "application/json")

	// Envoie la réponse JSON
	w.Write(jsonData)
}

func getAllSpotInfo(w http.ResponseWriter, r *http.Request) {
	// Simule la récupération des enregistrements depuis votre source de données (base de données, fichier, etc.)
	records := []SurfRecord{
		{
			ID: "rec5aF9TjMjBicXCK",
			Fields: RecordFields{
				SurfBreak:        []string{"Reef Break"},
				DifficultyLevel:  4,
				Destination:      "Pipeline",
				Geocode:          "🔵 eyJpIjoiUGlwZWxpbmUsIE9haHUsIEhhd2FpaSIsIm8iOnsic3RhdHVzIjoiT0siLCJmb3JtYXR0ZWRBZGRyZXNzIjoiRWh1a2FpIEJlYWNoIFBhcmssIEhhbGVpd2EsIEhJIDk2NzEyLCBVbml...",
				Influencers:      []string{"recD1zp1pQYc8O7l2", "rec1ptbRPxhS8rRun"},
				MagicSeaweedLink: "https://magicseaweed.com/Pipeline-Backdoor-Surf-Report/616/",
				Photos: []Photo{
					{
						ID:       "attf6yu03NAtCuv5L",
						URL:      "https://dl.airtable.com/ZuXJZ2NnTF40kCdBfTld_thomas-ashlock-64485-unsplash.jpg",
						Filename: "thomas-ashlock-64485-unsplash.jpg",
						Size:     688397,
						Type:     "image/jpeg",
						Thumbnails: Thumbnails{
							Small: Small{
								URL:    "https://dl.airtable.com/yfKxR9ZQqiT7drKxpjdF_small_thomas-ashlock-64485-unsplash.jpg",
								Width:  52,
								Height: 36,
							},
							Large: Large{
								URL:    "https://dl.airtable.com/cFfMuU8NQjaEskeC3B2h_large_thomas-ashlock-64485-unsplash.jpg",
								Width:  744,
								Height: 512,
							},
							Full: Full{
								URL:    "https://dl.airtable.com/psynuQNmSvOTe3BWa0Fw_full_thomas-ashlock-64485-unsplash.jpg",
								Width:  2233,
								Height: 1536,
							},
						},
					},
				},
				PeakSurfSeasonBegins:    "2018-07-22",
				DestinationStateCountry: "Oahu, Hawaii",
				PeakSurfSeasonEnds:      "2018-08-31",
				Address:                 "Pipeline, Oahu, Hawaii",
			},
			CreatedTime: "2018-05-31T00:16:16.000Z",
		},
		{
			ID: "recT98Z2El7YYwmc4",

			Fields: RecordFields{
				SurfBreak:        []string{"Point Break"},
				DifficultyLevel:  5,
				Destination:      "Skeleton Bay",
				Geocode:          "🔵 eyJpIjoiU2tlbGV0b24gQmF5LCBOYW1pYmlhIiwibyI6eyJzdGF0dXMiOiJPSyIsImZvcm1hdHRlZEFkZHJlc3MiOiJOYW1pYmlhIiwibGF0IjotMjUuOTE0NDkxOSwibG5nIjoxNC45MDY4NTk...",
				Influencers:      []string{"recD1zp1pQYc8O7l2", "rec1ptbRPxhS8rRun"},
				MagicSeaweedLink: "https://magicseaweed.com/Skeleton-Bay-Surf-Report/4591/",
				Photos: []Photo{
					{
						ID:       "attKinKKZvgdS5A5U",
						URL:      "https://dl.airtable.com/YzqA020RRLaTyAZAta9g_brandon-compagne-308937-unsplash.jpg",
						Filename: "brandon-compagne-308937-unsplash.jpg",
						Size:     1494974,
						Type:     "image/jpeg",
						Thumbnails: Thumbnails{
							Small: Small{
								URL:    "https://dl.airtable.com/Y8970kyrQHWfL6AMkxZQ_small_brandon-compagne-308937-unsplash.jpg",
								Width:  54,
								Height: 36,
							},
							Large: Large{
								URL:    "https://dl.airtable.com/dkQKXoUnTGiofIvg5TJR_large_brandon-compagne-308937-unsplash.jpg",
								Width:  768,
								Height: 512,
							},
							Full: Full{
								URL:    "https://dl.airtable.com/pexuxaQ6D2VV61pyhUwn_full_brandon-compagne-308937-unsplash.jpg",
								Width:  3000,
								Height: 2000,
							},
						},
					},
				},
				PeakSurfSeasonBegins:    "2018-09-01",
				DestinationStateCountry: "Namibia",
				PeakSurfSeasonEnds:      "2018-11-30",
				Address:                 "Skeleton Bay, Namibia",
			},
		},
		{
			ID: "recH2ennHFNOtB1Wt",
			Fields: RecordFields{
				SurfBreak:        []string{"Point Break"},
				DifficultyLevel:  4,
				Destination:      "Superbank",
				Geocode:          "🔵 eyJpIjoiU3VwZXJiYW5rLCBHb2xkIENvYXN0LCBBdXN0cmFsaWEiLCJvIjp7InN0YXR1cyI6Ik9LIiwiZm9ybWF0dGVkQWRkcmVzcyI6IlNuYXBwZXIgUm9ja3MgUmQsIENvb2xhbmdhdHRhIFF...",
				Influencers:      []string{"recSkJ4HuvzAUBrdd"},
				MagicSeaweedLink: "https://magicseaweed.com/Surfers-Paradise-Gold-Coast-Surf-Report/1012/",
				Photos: []Photo{
					{
						ID:       "attmtbEOAQteRjz2p",
						URL:      "https://dl.airtable.com/I4E4xZeQbO2g814udQDm_jeremy-bishop-80371-unsplash.jpg",
						Filename: "jeremy-bishop-80371-unsplash.jpg",
						Size:     1524876,
						Type:     "image/jpeg",
						Thumbnails: Thumbnails{
							Small: Small{
								URL:    "https://dl.airtable.com/IWP9RPvvSM2pX1sHeigV_small_jeremy-bishop-80371-unsplash.jpg",
								Width:  48,
								Height: 36,
							},
							Large: Large{
								URL:    "https://dl.airtable.com/aBnINo8qQqDvER2f2wGg_large_jeremy-bishop-80371-unsplash.jpg",
								Width:  683,
								Height: 512,
							},
							Full: Full{
								URL:    "https://dl.airtable.com/eZxgLD8hRI2Y27J39LNl_full_jeremy-bishop-80371-unsplash.jpg",
								Width:  3000,
								Height: 2250,
							},
						},
					},
				},
				PeakSurfSeasonBegins:    "2018-11-28",
				DestinationStateCountry: "Gold Coast, Australia",
				PeakSurfSeasonEnds:      "2019-02-18",
			},
			CreatedTime: "2018-05-31T00:16:16.000Z",
		},
	}

	// Convertit les enregistrements en JSON
	jsonData, err := json.Marshal(records)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}

	// Définit le type de contenu de la réponse HTTP comme application/json
	w.Header().Set("Content-Type", "application/json")

	// Envoie la réponse JSON
	w.Write(jsonData)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/spots", getSpotList).Methods("GET")
	router.HandleFunc("/spotsAll", getAllSpotInfo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
