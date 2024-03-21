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

func (user *UserModel) toDto() map[string]string {
	return map[string]string{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"role":     user.Role,
	}
}

func (user *UserModel) createCacheKey(id int) string {
	return fmt.Sprintf("users_%d", id)
}

func (user *UserModel) AddUser() int {
	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	user_map := user.toDto()

	id := connector.InsertIntoTable("users", user_map)

	utility.ApplicationCache.AddElement(user.createCacheKey(id), user_map)

	return id
}

func (user *UserModel) GetUserById(id int) UserModel {
	userMap, err := utility.ApplicationCache.GetElement(user.createCacheKey(id))
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

func (user *UserModel) ModifyUserById(id int) int {
	utility.ApplicationCache.RemoveElement(user.createCacheKey(id))

	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	where_clause := map[string]string{
		"id": strconv.Itoa(id),
	}

	num_rows, err := connector.UpdateTableWhere("users", where_clause, user.toDto())

	if err != nil {
		panic(err)
	}

	user_map := user.toDto()
	utility.ApplicationCache.AddElement(user.createCacheKey(id), user_map)

	return num_rows

}

func (user *UserModel) DeleteUserById(id int) int {
	utility.ApplicationCache.RemoveElement(user.createCacheKey(id))

	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	where_clause := map[string]string{
		"id": strconv.Itoa(id),
	}

	num_rows, err := connector.DeleteFromTableWhere("users", where_clause)
	if err != nil {
		panic(err)
	}
	return num_rows
}

func (user *UserModel) SearchUser() UserModel {
	connector := utility.NewDatabasePostgreSQLConnector()
	connector.OpenConnection()
	defer connector.CloseConnection()

	where_clause := map[string]string{
		"password": user.Password,
	}
	if user.Username != "" {
		where_clause["username"] = user.Username
	}
	if user.Email != "" {
		where_clause["email"] = user.Email
	}
	userMap, err := connector.SelectFromTableWhere("users", where_clause)
	if err != nil {
		panic(err)
	}
	ret_value, _ := user.fromDto(userMap)
	return ret_value
}
