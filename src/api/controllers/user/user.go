package user

import (
	"strconv"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	common_function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	sql_event "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/controllers/user/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.User
	)

	if err = c.ShouldBindJSON(&request); err != nil {
		return c, nil, err
	}

	hashedPassword, err := common_function.HashPassword(*request.Password)
	if err != nil {
		return
	}

	request.NormalizedUser()
	request.Password = &hashedPassword

	if err = sql_event.CreateUser(db, request); err != nil {
		return c, &constants.InsertSuccess, err
	} else {
		return c, &err, nil
	}
}

func Delete(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		request models.User
	)

	UserIDstr := c.Query("id")
	UserID, err := strconv.Atoi(UserIDstr)
	if err != nil {
		return c, nil, err
	}

	request.ID = uint(UserID)
	if err = sql_event.Delete(db, request); err != nil {
		return c, nil, err
	} else {
		return c, &constants.DeleteSuccess, nil
	}
}

func Get(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		users *[]models.User
	)

	if UserIDstr := c.Query("id"); UserIDstr != "" {
		if users, err := sql_event.GetByParam(db, "id", UserIDstr); err != nil {
			return c, &err, nil
		} else {
			return c, users, err
		}
	}
	if users, err = sql_event.Get(db); err != nil {
		return c, &err, nil
	}
	return c, users, err

}

func ResetPassword(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		user *models.User
	)

	if UserIDstr := c.Query("id"); UserIDstr != "" {

		UserID, err := strconv.Atoi(UserIDstr)
		if err != nil {
			return c, nil, err
		}

		if UserIDstr := c.Query("id"); UserIDstr != "" {
			if users, err := sql_event.GetByParam(db, "id", UserIDstr); err != nil {
				return c, &err, nil
			} else {
				return c, users, err
			}
		}
		user = &models.User{}
		user.ID = uint(UserID)
		if _, err := sql_event.PutDefaultPassword(db, user); err != nil {
			return c, &err, nil
		} else {
			return c, &constants.UpdateSuccess, err
		}
	}

	return c, &err, nil

}

func PutBy(c *gin.Context, db *gorm.DB) (context *gin.Context, data interface{}, err error) {
	var (
		user *models.User
	)

	if UserIDstr := c.Query("id"); UserIDstr != "" {

		UserID, err := strconv.Atoi(UserIDstr)
		if err != nil {
			return c, nil, err
		}

		if UserIDstr := c.Query("id"); UserIDstr != "" {
			var field, value string
			if UserName := c.Query("name"); UserName != "" {
				field = "name"
				value = UserName
			}

			if UserLastName := c.Query("LastName"); UserLastName != "" {
				field = "last_name"
				value = UserLastName
			}

			if UserEmail := c.Query("email"); UserEmail != "" {
				field = "email"
				value = UserEmail
			}

			if UserType := c.Query("Type"); UserType != "" {
				field = "type"
				value = UserType
			}

			user = &models.User{}
			user.ID = uint(UserID)
			if err := sql_event.PutBy(db, field, value, user); err != nil {
				return c, &err, nil
			} else {
				return c, &constants.UpdateSuccess, err
			}
		}
	}

	return c, &err, nil

}
