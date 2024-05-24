package main

import (
	"database/sql"
	"log"
)

func queryHeroes(db *sql.DB) {
	// Query the Heroes table
	rows, err := db.Query("SELECT ID, Name, Universe, Skill FROM Heroes")
	if err != nil {
		log.Fatalf("Error querying Heroes table: %v", err)
	}
	defer rows.Close()

	// Iterate over the rows and print the results
	for rows.Next() {
		var (
			id       int
			name     string
			universe string
			skill    string
		)
		if err := rows.Scan(&id, &name, &universe, &skill); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		log.Printf("Hero: ID=%d, Name=%s, Universe=%s, Skill=%s", id, name, universe, skill)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}
}
