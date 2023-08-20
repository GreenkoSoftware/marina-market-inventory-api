package sql

import (
	"context"
	"time"

	constants "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/constant"
	common_function "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/function"
	"github.com/GreenkoSoftware/marina-market-inventory-api/src/api/common/sql"
	models "github.com/GreenkoSoftware/marina-market-inventory-api/src/api/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user models.User) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Create(&user)

	if results.Error != nil {
		return results.Error
	}

	return nil
}

func Delete(db *gorm.DB, user models.User) (err error) {
	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		WithContext(ctx).
		Delete(&user)

	if results.Error != nil {
		return results.Error
	}

	return nil
}

func Get(db *gorm.DB) (users *[]models.User, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Omit("password").
		Preload("UserType").
		WithContext(ctx).
		Find(&users).Error

	if results != nil {
		return nil, results
	}

	return users, nil
}

func GetByParam(db *gorm.DB, fiel string, value string) (users *[]models.User, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Omit("password").
		Scopes(sql.By(fiel, value)).
		Preload("UserType").
		WithContext(ctx).
		Find(&users).Error

	if results != nil {
		return nil, results
	}

	return users, nil
}

func GetUser(db *gorm.DB, userEmail string) (user *models.User, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Preload("UserType").
		WithContext(ctx).
		Where("email = ?", userEmail).
		Find(&user).Error

	if results != nil || user.Password == nil {
		return nil, constants.ErrorInPassword
	}

	return user, nil
}

func PutDefaultPassword(db *gorm.DB, user *models.User) (users *[]models.User, err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	hashedPassword, err := common_function.HashPassword(constants.DefaultPass)
	if err != nil {
		return
	}
	results := db.
		WithContext(ctx).
		Model(&user).
		Update("password", hashedPassword).Error

	if results != nil {
		return nil, results
	}

	return users, nil
}

func PutBy(db *gorm.DB, field string, value string, user *models.User) (err error) {

	var ctx, cancel = context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	results := db.
		Omit("password").
		Preload("UserType").
		WithContext(ctx).
		Model(&user).
		Update(field, value).Error

	if results != nil {
		return results
	}

	return nil
}
