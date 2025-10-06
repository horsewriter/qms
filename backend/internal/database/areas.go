package database

import (
	"quality-system/models"
)

func (db *DB) GetAreas() ([]models.Area, error) {
	var areas []models.Area
	rows, err := db.Query("SELECT * FROM areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var area models.Area
		if err := rows.Scan(&area.ID, &area.Name); err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}

func (db *DB) GetAreaByID(id int) (*models.Area, error) {
	var area models.Area
	err := db.QueryRow("SELECT * FROM areas WHERE id = ?", id).Scan(&area.ID, &area.Name)
	if err != nil {
		return nil, err
	}
	return &area, nil
}

func (db *DB) CreateArea(area models.Area) (int64, error) {
	result, err := db.Exec("INSERT INTO areas (name) VALUES (?)", area.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) UpdateArea(area models.Area) error {
	_, err := db.Exec("UPDATE areas SET name = ? WHERE id = ?", area.Name, area.ID)
	return err
}

func (db *DB) DeleteArea(id int) error {
	_, err := db.Exec("DELETE FROM areas WHERE id = ?", id)
	return err
}

func (db *DB) SearchAreas(search string) ([]models.Area, error) {
	var areas []models.Area
	rows, err := db.Query("SELECT * FROM areas WHERE name LIKE ?", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var area models.Area
		if err := rows.Scan(&area.ID, &area.Name); err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}
