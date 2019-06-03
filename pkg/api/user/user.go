// Package user contains user application services
package user

import (
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// Query returns list of users
func (u *User) Query(c echo.Context, req *schemago.ReqUser) ([]schemago.RespUser, int, error) {
	return u.udb.Query(u.db, req)
}

// Get returns single user
func (u *User) Get(c echo.Context, userid string) (*schemago.RespUser, error) {
	return u.udb.Get(u.db, userid)
}

// Create creates a new user account
func (u *User) Create(c echo.Context, req *schemago.ReqCreateUser) (*schemago.SUser, error) {
	// if err := u.rbac.AccountCreate(c, req.RoleID, req.CompanyID, req.LocationID); err != nil {
	// 	return nil, err
	// }
	req.Password = u.sec.Hash(req.Password)
	return u.udb.Create(u.db, u.ce, req)
}

// Delete deletes a user
func (u *User) Delete(c echo.Context, userid string) error {
	return u.udb.Delete(u.db, u.ce, userid)
}

// Update updates user's contact information
func (u *User) Update(c echo.Context, req *schemago.ReqUpdateUser) (*schemago.SUser, error) {

	user, err := u.udb.Update(u.db, u.ce, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Enable updates user's contact information
func (u *User) Enable(c echo.Context, userid string, active bool) error {
	return u.udb.Enable(u.db, u.ce, userid, active)
}

// Disable updates user's contact information
func (u *User) Disable(c echo.Context, userid string, active bool) error {
	return u.udb.Disable(u.db, u.ce, userid, active)
}
