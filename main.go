package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"example.com/mod/entity"
)

func main() {
	file, err := os.Open("/Users/flodht/Documents/Dev/Formula-Calendar/ressources/data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var grandPrix entity.GrandPrix
	json.Unmarshal(byteValue, &grandPrix)

	fmt.Println("Grand Prix Data:", grandPrix)

	// Modifier les données
	grandPrix.Name = "Updated Grand Prix Name"

	// Écrire les données dans un fichier JSON
	newData, _ := json.MarshalIndent(grandPrix, "", "  ")
	err = os.WriteFile("/Users/flodht/Documents/Dev/Formula-Calendar/ressources/data.json", newData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
