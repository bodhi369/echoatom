package schemago

import "time"
import "github.com/jinzhu/gorm"

// SUser 用户
type SUser struct {
	gorm.Model
	Recid    string `json:"recid" gorm:"type:varchar(40);not null;unique"`
	Code     string `json:"code" gorm:"type:varchar(50);not null;unique"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	NickName string `json:"nick_name" gorm:"type:varchar(50)"`

	Password string `json:"-"`
	Email    string `json:"email"`

	Mobile  string `json:"mobile"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Active bool `json:"active"`

	LastLogin          time.Time `json:"last_login" `
	LastPasswordChange time.Time `json:"last_password_change"`

	Token string `json:"-"`
}

// ReqCreateRole ReqCreateRole
type ReqCreateRole struct {
	Roleid string `json:"roleid" gorm:"type:varchar(40)"`
}

// ReqCreateRoles ReqCreateRoles
type ReqCreateRoles []ReqCreateRole

// ReqCreateUser ReqCreateUser
type ReqCreateUser struct {
	Recid    string `json:"recid" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,alphanum"`
	NickName string `json:"nick_name" `

	Password        string `json:"password" validate:"required,min=4"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=4"`

	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Active bool           `json:"active" validate:"required"`
	Roles  ReqCreateRoles `json:"roles" validate:"required"`
}

// ReqUpdateUser ReqUpdateUser
type ReqUpdateUser struct {
	Recid    string `json:"recid" validate:"required"`
	Code     string `json:"code" validate:"required"`
	Name     string `json:"name" validate:"required,min=3,alphanum"`
	NickName string `json:"nick_name" `

	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	Active bool           `json:"active" validate:"required"`
	Roles  ReqCreateRoles `json:"roles" validate:"required"`
}

// ToSchemaSUserRole ToSchemaSUserRole
func (c *ReqCreateRole) ToSchemaSUserRole(userid string) *SUserRole {
	return &SUserRole{
		Userid: userid,
		Roleid: c.Roleid,
	}
}

// ToSchemaSUser ToSchemaSUser
func (c *ReqCreateUser) ToSchemaSUser() *SUser {
	return &SUser{
		Recid:    c.Recid,
		Code:     c.Code,
		Name:     c.NickName,
		Password: c.Password,
		Email:    c.Email,
		Mobile:   c.Mobile,
		Phone:    c.Phone,
		Address:  c.Address,
		Active:   c.Active,
	}
}

// ToSchemaSUserRoles ToSchemaSUserRoles
func (c *ReqCreateUser) ToSchemaSUserRoles() []*SUserRole {
	userRoles := make([]*SUserRole, len(c.Roles))

	for i, item := range c.Roles {
		userRoles[i] = item.ToSchemaSUserRole(c.Recid)
	}
	return userRoles
}

// ToSchemaSUser ToSchemaSUser
func (c *ReqUpdateUser) ToSchemaSUser() *SUser {
	return &SUser{
		Recid:   c.Recid,
		Code:    c.Code,
		Name:    c.NickName,
		Email:   c.Email,
		Mobile:  c.Mobile,
		Phone:   c.Phone,
		Address: c.Address,
		Active:  c.Active,
	}
}

// ToSchemaSUserRoles ToSchemaSUserRoles
func (c *ReqUpdateUser) ToSchemaSUserRoles() []*SUserRole {
	userRoles := make([]*SUserRole, len(c.Roles))

	for i, item := range c.Roles {
		userRoles[i] = item.ToSchemaSUserRole(c.Recid)
	}
	return userRoles
}

// ChangePassword updates user's password related fields
func (u *SUser) ChangePassword(hash string) {
	u.Password = hash
	u.LastPasswordChange = time.Now()
}

// UpdateLastLogin updates last login field
func (u *SUser) UpdateLastLogin(token string) {
	u.Token = token
	u.LastLogin = time.Now()
}

// SUserRole 用户对应的角色
type SUserRole struct {
	gorm.Model
	Userid string `json:"userid" gorm:"type:varchar(40)"`
	Roleid string `json:"roleid" gorm:"type:varchar(40)"`
}

// SUserRoles SUserRole
type SUserRoles []SUserRole

// RespUser RespUser
type RespUser struct {
	SUser
	SUserRoles
}

// ReqUser  request user
type ReqUser struct {
	PaginationReq
	Code string `json:"code" validate:"omitempty"`
	Name string `json:"name" validate:"omitempty"`
}
