package model

import (
	"database/sql"
	"fmt"

	"CoachingFireBackend/internal/utility"

	_ "github.com/lib/pq"
)

type UserModel struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUserModel() *UserModel {
	user_model := &UserModel{}

	return user_model
}

func openConnection(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}

func (user *UserModel) AddUser(username string, email string, password string, role utility.Role) int {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", utility.Configuration.Database.User, utility.Configuration.Database.Password, utility.Configuration.Database.Name)
	db := openConnection(connStr)

	var id int
	err := db.QueryRow("INSERT INTO users(username, email, role, password_hash) VALUES($1,$2,$3,$4) RETURNING id", username, email, role, password).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
