package mssql

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// 定义错误
var (
	ErrNotExists = echo.NewHTTPError(http.StatusInternalServerError, "user code not exists.")
)

// FindByCode queries for single user by username
func (u *User) FindByCode(db *gorm.DB, code string) (*schemago.SUser, error) {
	var user = new(schemago.SUser)
	if db.Find(&user, "code = ?", code).RecordNotFound() {
		return nil, ErrNotExists
	}
	return user, nil
}

// FindByToken queries for single user by token
func (u *User) FindByToken(db *gorm.DB, token string) (*schemago.SUser, error) {
	var user = new(schemago.SUser)
	if db.Where("token = ?", token).First(&user).Find(&user).RecordNotFound() {
		return nil, ErrNotExists
	}
	return user, nil
}

// Update updates user's info
func (u *User) Update(db *gorm.DB, user *schemago.SUser) error {
	return db.Save(&user).Error
}
