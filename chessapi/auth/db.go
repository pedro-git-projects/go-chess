package auth

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := "postgres://postgres@localhost/chess?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func loadUsersFromDB(db *Database) ([]User, error) {
	query := "SELECT username, password FROM users"
	rows, err := db.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
