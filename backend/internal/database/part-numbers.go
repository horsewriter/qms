package database

import (
	"quality-system/models"
)

func (db *DB) GetPartNumbers() ([]models.PartNumber, error) {
	var partNumbers []models.PartNumber
	rows, err := db.Query("SELECT * FROM part_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var partNumber models.PartNumber
		if err := rows.Scan(&partNumber.ID, &partNumber.Number); err != nil {
			return nil, err
		}
		partNumbers = append(partNumbers, partNumber)
	}
	return partNumbers, nil
}

func (db *DB) GetPartNumberByID(id int) (*models.PartNumber, error) {
	var partNumber models.PartNumber
	err := db.QueryRow("SELECT * FROM part_numbers WHERE id = ?", id).Scan(&partNumber.ID, &partNumber.Number)
	if err != nil {
		return nil, err
	}
	return &partNumber, nil
}

func (db *DB) CreatePartNumber(partNumber models.PartNumber) (int64, error) {
	result, err := db.Exec("INSERT INTO part_numbers (number) VALUES (?)", partNumber.Number)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) UpdatePartNumber(partNumber models.PartNumber) error {
	_, err := db.Exec("UPDATE part_numbers SET number = ? WHERE id = ?", partNumber.Number, partNumber.ID)
	return err
}

func (db *DB) DeletePartNumber(id int) error {
	_, err := db.Exec("DELETE FROM part_numbers WHERE id = ?", id)
	return err
}

func (db *DB) SearchPartNumbers(search string) ([]models.PartNumber, error) {
	var partNumbers []models.PartNumber
	rows, err := db.Query("SELECT * FROM part_numbers WHERE number LIKE ?", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var partNumber models.PartNumber
		if err := rows.Scan(&partNumber.ID, &partNumber.Number); err != nil {
			return nil, err
		}
		partNumbers = append(partNumbers, partNumber)
	}
	return partNumbers, nil
}
