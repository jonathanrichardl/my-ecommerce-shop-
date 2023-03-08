package repository

import (
	"database/sql"
	"my-ecommerce-shop/order-management/internal/repository/model"

	_ "github.com/lib/pq"
)

type UserRepository interface {
	Get(id string) (*model.User, error)
	Create(user model.User) (string, error)
}

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImplementation {
	return &UserRepositoryImplementation{db: db}
}

func (r *UserRepositoryImplementation) Get(id string) (*model.User, error) {
	stmt, err := r.db.Prepare(`SELECT id, username, password, email, phone_number, role FROM users WHERE id = $1`)
	if err != nil {
		return nil, err
	}

	result := stmt.QueryRow(id)

	var user model.User
	err = result.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImplementation) Create(user model.User) (string, error) {
	stmt, err := r.db.Prepare(`INSERT INTO users 
	(id, username, email, password, phone_number, role) 
	VALUES ($1,$2,$3,$4,$5,$6)`)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(user.Id, user.Username, user.Email, user.Password, user.PhoneNumber, user.Role)

	return user.Id, err
}
