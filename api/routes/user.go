package routes

import (
	"CoachingFireBackend/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
}

func NewUser() *UserRoutes {
	user := &UserRoutes{}
	return user
}

func (user *UserRoutes) CreateUserHandler(c *gin.Context) {
	var newUser model.UserModel
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser.AddUser())
}

func (user *UserRoutes) GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Aggiungi la logica di validazione e gestione dell'utente
	userModel := model.NewUserModel()
	id_num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	c.JSON(http.StatusOK, userModel.GetUserById(id_num))
}

func (user *UserRoutes) ModifyUserById(c *gin.Context) {
	id := c.Param("id")

	var updateUser model.UserModel
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id_num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	c.JSON(http.StatusOK, updateUser.ModifyUserById(id_num))

}

func (user *UserRoutes) DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	userModel := model.NewUserModel()
	id_num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	c.JSON(http.StatusOK, userModel.DeleteUserById(id_num))
}
