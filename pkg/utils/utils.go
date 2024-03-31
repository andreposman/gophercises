package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

func PrintName(number, name string) {
	fmt.Println("--------------------------------------")
	fmt.Printf("Gophercises | %v - %v \n", number, name)
	fmt.Println("--------------------------------------")
}

// ReadCSV The path is in relation to the main.go file in the root of the project.
func ReadCSV(filePath, fileName string) ([][]string, error) {
	log.Info().Msgf("Reading: %v%v", filePath, fileName)
	path := filepath.Join(filePath, fileName)

	file, err := os.Open(path)
	if err != nil {
		log.Error().Msgf("Error opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Error().Msgf("Error reading csv file: %v", err)
		return nil, err
	}

	return csvData, nil
}
