package auth

import (
	"os"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	common_function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	user_sql "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/user/sql"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/middlewares"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {

	var (
		request models.User
		userBd  *models.User
	)

	// Get the request body from context and validate it against model struct
	if err = c.ShouldBindJSON(&request); err != nil {
		return c, nil, err
	}

	request.NormalizedUser()
	//get user from db by email
	if userBd, err = user_sql.GetUser(db, request.Email); err != nil {
		return c, nil, err
	}

	// validate if hashed password is the same in database
	err = common_function.ComparePasswords(*userBd.Password, *request.Password)
	if err != nil {
		return c, nil, constants.ErrorInPassword
	}

	userBd.FailedAttempts = 0
	userBd.Locked = false

	// create a jwf
	token, err := middlewares.CreateToken([]byte(os.Getenv("JWT")), request.Email)
	if err != nil {
		return
	}

	userBd.Password = nil
	// save auth response to send to client
	response := &models.AuthResponse{
		Token:      token,
		UserType:   userBd.UserType.TypeName,
		UserTypeID: userBd.UserTypeID,
		User:       *userBd,
	}

	return c, response, nil
}
