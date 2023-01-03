package handler

import (
	"net/http"
	"project-go/models/user"
	"project-go/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	userService user.Service
}

func MigrateAndGetUser(db *gorm.DB) (userHandler *userHandler) {
	db.AutoMigrate(&user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler = NewUserHandler(userService)
	return
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Register godoc
//
//	@Summary	Register a user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		userInput	body		user.InputUser	true	"Message body"
//	@Success	201			{object}	user.User
//	@Failure	400			{object}	Response
//	@Failure	401			{object}	Response
//	@Router		/public/user/register [post]
func (th *userHandler) Register(c *gin.Context) {
	var input user.InputUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := th.userService.Register(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Error: Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "New user registered",
		Data:    newUser,
	}
	c.JSON(http.StatusCreated, response)
}

// Login godoc
//
//	@Summary	Login a user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Param		userInput	body		user.InputUser	true	"Message body"
//	@Success	200			{object}	Response
//	@Failure	400			{object}	Response
//	@Failure	401			{object}	Response
//	@Router		/public/user/login [post]
func (th *userHandler) Login(c *gin.Context) {
	var input user.InputUser
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tkn, err := th.userService.Login(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Login error",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Success: true,
		Message: "Login success",
		Data:    token.Token{Token: tkn},
	}
	c.JSON(http.StatusCreated, response)
}

// Current godoc
//
//	@Summary	Get current user
//	@Tags		user
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	Response
//	@Failure	400	{object}	Response
//	@Failure	401	{object}	Response
//	@Router		/protected/user/current [get]
//
//	@Security	BearerAuth
func (th *userHandler) CurrentUser(c *gin.Context) {

	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := th.userService.GetById(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
