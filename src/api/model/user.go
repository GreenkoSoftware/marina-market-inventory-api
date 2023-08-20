package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string    `json:"name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password       *string   `json:"password,omitempty" binding:"required"`
	LastAttempt    time.Time `json:"-"`
	FailedAttempts int       `json:"failed_attempts"`
	Locked         bool      `json:"-"`
	UserTypeID     int       `json:"-"`
	UserType       UserType  `gorm:"foreignKey:UserTypeID"`
}

type UserType struct {
	gorm.Model
	TypeName string `gorm:"unique;not null" json:"type_name"`
}

func (user *User) NormalizedUser() {
	user.Name = strings.ToLower(user.Name)
	user.LastName = strings.ToLower(user.LastName)
	user.Email = strings.ToLower(user.Email)
	user.UserTypeID = 2
}
