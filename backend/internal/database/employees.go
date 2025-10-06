package database

import "quality-system/models"

func (db *DB) GetEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	rows, err := db.Query("SELECT id, name, role FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Role); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (db *DB) GetEmployeeByID(id int) (*models.Employee, error) {
	var employee models.Employee
	err := db.QueryRow("SELECT id, name, role FROM employees WHERE id = ?", id).Scan(&employee.ID, &employee.Name, &employee.Role)
	return &employee, err
}

func (db *DB) CreateEmployee(employee models.Employee) (int64, error) {
	result, err := db.Exec("INSERT INTO employees (name, role) VALUES (?, ?)", employee.Name, employee.Role)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (db *DB) UpdateEmployee(employee models.Employee) error {
	_, err := db.Exec("UPDATE employees SET name = ?, role = ? WHERE id = ?", employee.Name, employee.Role, employee.ID)
	return err
}

func (db *DB) DeleteEmployee(id int) error {
	_, err := db.Exec("DELETE FROM employees WHERE id = ?", id)
	return err
}

func (db *DB) SearchEmployees(search string) ([]models.Employee, error) {
	var employees []models.Employee
	rows, err := db.Query("SELECT id, name, role FROM employees WHERE name LIKE ? OR role LIKE ?", "%"+search+" sprinting%", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Role); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
