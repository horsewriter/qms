package database

import (
	"quality-system/models"
)

func (db *DB) GetLevels() ([]models.Level, error) {
	var levels []models.Level
	rows, err := db.Query("SELECT * FROM levels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var level models.Level
		if err := rows.Scan(&level.ID, &level.Name); err != nil {
			return nil, err
		}
		levels = append(levels, level)
	}
	return levels, nil
}

func (db *DB) GetLevelByID(id int) (*models.Level, error) {
	var level models.Level
	err := db.QueryRow("SELECT * FROM levels WHERE id = ?", id).Scan(&level.ID, &level.Name)
	if err != nil {
		return nil, err
	}
	return &level, nil
}

func (db *DB) CreateLevel(level models.Level) (int64, error) {
	result, err := db.Exec("INSERT INTO levels (name) VALUES (?)", level.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) UpdateLevel(level models.Level) error {
	_, err := db.Exec("UPDATE levels SET name = ? WHERE id = ?", level.Name, level.ID)
	return err
}

func (db *DB) DeleteLevel(id int) error {
	_, err := db.Exec("DELETE FROM levels WHERE id = ?", id)
	return err
}

func (db *DB) SearchLevels(search string) ([]models.Level, error) {
	var levels []models.Level
	rows, err := db.Query("SELECT * FROM levels WHERE name LIKE ?", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var level models.Level
		if err := rows.Scan(&level.ID, &level.Name); err != nil {
			return nil, err
		}
		levels = append(levels, level)
	}
	return levels, nil
}
