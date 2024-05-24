package main

import (
	"NGC_Avenger/config"
	"NGC_Avenger/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func heroesHandler(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Universe, Skill FROM heroes")
	if err != nil {
		http.Error(w, "Failed to query heroes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []entity.Heroes

	for rows.Next() {
		var hero entity.Heroes
		if err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill); err != nil {
			http.Error(w, "Failed to scan hero", http.StatusInternalServerError)
			return
		}
		heroes = append(heroes, hero)
	}
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Encode the heroes slice directly to the HTTP response
	if err := json.NewEncoder(w).Encode(heroes); err != nil {
		http.Error(w, "Failed to encode heroes to JSON", http.StatusInternalServerError)
		return
	}

}

func villainsHandler(w http.ResponseWriter, r *http.Request) {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Universe FROM villains")
	if err != nil {
		http.Error(w, "Failed to query villains", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []entity.Villains

	for rows.Next() {
		var villain entity.Villains
		if err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe); err != nil {
			http.Error(w, "Failed to scan villain", http.StatusInternalServerError)
			return
		}
		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(villains); err != nil {
		http.Error(w, "Failed to encode villains to JSON", http.StatusInternalServerError)
		return
	}
}

func testHeroes() {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Universe, Skill FROM heroes")
	if err != nil {
		fmt.Println("Error querying", err)
		return
	}
	defer rows.Close()

	var heroes []entity.Heroes

	for rows.Next() {
		var hero entity.Heroes
		if err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill); err != nil {
			fmt.Println("Error Scanning", err)
			return
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
}

func testVillains() {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT ID, Name, Universe FROM villains")
	if err != nil {
		fmt.Println("Error querying", err)
		return
	}
	defer rows.Close()

	var Villains []entity.Villains

	for rows.Next() {
		var villains entity.Villains
		if err := rows.Scan(&villains.ID, &villains.Name, &villains.Universe); err != nil {
			fmt.Println("Error Scanning", err)
			return
		}
		Villains = append(Villains, villains)
	}
	fmt.Println(Villains)
}

// // heroesHandler handles the /heroes endpoint
// func heroesHandler(w http.ResponseWriter, r *http.Request) {

// 	db := config.ConnectDB()
// 	defer db.Close()

// 	rows, err := db.Query("SELECT ID, Name, Universe, Skill FROM heroes")
// 	if err != nil {
// 		http.Error(w, "Failed to query heroes", http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var heroes []entity.Heroes
// 	for rows.Next() {
// 		var hero entity.Heroes
// 		if err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill); err != nil {
// 			http.Error(w, "Failed to scan hero", http.StatusInternalServerError)
// 			return
// 		}
// 		heroes = append(heroes, hero)
// 	}
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, "Failed to iterate over rows", http.StatusInternalServerError)
// 		return
// 	}

// 	// Convert the heroes data to JSON
// 	w.Header().Set("Content-Type", "application/json")
// 	if err := json.NewEncoder(w).Encode(heroes); err != nil {
// 		http.Error(w, "Failed to encode heroes to JSON", http.StatusInternalServerError)
// 		return
// 	}
// }

func main() {
	// Create a new HTTP server
	// http.HandleFunc("/heroes", heroesHandler)
	// fmt.Println("Starting server on :8000...")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	http.HandleFunc("/heroes", heroesHandler)
	http.HandleFunc("/villains", villainsHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)

}
