package repository

import (
	"database/sql"
	"example.com/enterprise_grade_go_project/pkg/model"
)

// UserRepository represents the user repository.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository returns a new instance of the UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUsers returns a list of users.
func (u *UserRepository) GetUsers() ([]*model.User, error) {
	// Query the database
	rows, err := u.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Scan the rows
	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}