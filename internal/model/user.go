package model

import (
	"CoachingFireBackend/internal/utility"
	"fmt"
	"strconv"

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

func (user *UserModel) fromDto(userMap interface{}) (UserModel, error) {
	if userMap, ok := userMap.(map[string]string); ok {
		return UserModel{
			Username: userMap["username"],
			Email:    userMap["email"],
			Password: userMap["password"],
			Role:     userMap["role"],
		}, nil
	}
	return UserModel{}, fmt.Errorf("Map is not a user")
}

func (user *UserModel) AddUser(username string, email string, password string, role utility.Role) int {
	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	user_map := map[string]string{
		"username": username,
		"email":    email,
		"role":     utility.GetRoleString(role),
		"password": password,
	}

	id := connector.InsertIntoTable("users", user_map)

	utility.ApplicationCache.AddElement(fmt.Sprintf("users_%d", id), user_map)

	return id
}

func (user *UserModel) GetUserById(id int) UserModel {
	userMap, err := utility.ApplicationCache.GetElement(fmt.Sprintf("users_%d", id))
	if err == nil {
		ret_value, _ := user.fromDto(userMap)
		return ret_value
	}

	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	where_clause := map[string]string{
		"id": strconv.Itoa(id),
	}
	userMap, err = connector.SelectFromTableWhere("users", where_clause)
	if err != nil {
		panic(err)
	}
	ret_value, _ := user.fromDto(userMap)
	return ret_value
}
