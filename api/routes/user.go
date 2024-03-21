package routes

import (
	"CoachingFireBackend/internal/model"
	"CoachingFireBackend/internal/utility"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
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

	// Aggiungi la logica di validazione e gestione dell'utente
	user_model := model.NewUserModel()
	user_id := user_model.AddUser(newUser.Username, newUser.Email, newUser.Password, utility.Role(newUser.Role))

	c.JSON(http.StatusCreated, user_id)
}

func (user *UserRoutes) GetUserById(c *gin.Context) {
	id := c.Param("id")

	// Aggiungi la logica di validazione e gestione dell'utente
	user_model := model.NewUserModel()
	id_num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error during conversion:", err)
		return
	}

	c.JSON(http.StatusOK, user_model.GetUserById(id_num))
}
