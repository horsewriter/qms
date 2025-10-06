package utils

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"quality-system/internal/database"
	"quality-system/internal/models"
)

func LoadCustomersFromCSV(db *database.DB, csvPath string) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 2 {
			continue
		}

		customerID := strings.TrimSpace(record[0])
		customerName := strings.TrimSpace(record[1])

		if customerID == "" || customerName == "" {
			continue
		}

		customer := models.Customer{
			CustomerID: customerID,
			Name:       customerName,
		}

		_, err := db.CreateCustomer(ctx, customer)
		if err != nil {
			log.Printf("Error creating customer %s: %v", customerName, err)
		}
	}

	log.Printf("Loaded customers from %s", csvPath)
	return nil
}

func LoadPartNumbersFromCSV(db *database.DB, csvPath string) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()

	for i, record := range records {
		if i == 0 {
			continue
		}
		if len(record) < 2 {
			continue
		}

		partID := strings.TrimSpace(record[0])
		partName := strings.TrimSpace(record[1])

		if partID == "" || partName == "" {
			continue
		}

		partNumber := models.PartNumber{
			Number:     partName,
			Customer:   "",
			CustomerID: "",
		}

		_, err := db.CreatePartNumber(ctx, partNumber)
		if err != nil {
			log.Printf("Error creating part number %s: %v", partName, err)
		}
	}

	log.Printf("Loaded part numbers from %s", csvPath)
	return nil
}
