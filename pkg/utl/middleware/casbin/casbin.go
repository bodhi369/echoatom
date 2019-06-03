package casbinmw

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New generates new Casbin service necessery for auth middleware
func New(ce *casbin.Enforcer) *Casbinauth {
	return &Casbinauth{
		Skipper:  middleware.DefaultSkipper,
		Enforcer: ce,
	}
}

// Casbinauth defines the config for Casbinauth middleware.
type Casbinauth struct {
	// Skipper defines a function to skip middleware.
	Skipper middleware.Skipper

	// Enforcer Casbinauth main rule.
	// Required.
	Enforcer *casbin.Enforcer
}

// MWFunc returns a Casbinauth middleware.
//
// For valid credentials it calls the next handler.
// For missing or invalid credentials, it sends "401 - Unauthorized" response.
func (a *Casbinauth) MWFunc() echo.MiddlewareFunc {
	if a.Skipper == nil {
		a.Skipper = middleware.DefaultSkipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if a.Skipper(c) || a.CheckPermission(c) {
				return next(c)
			}

			return echo.ErrForbidden
		}
	}
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *Casbinauth) GetUserName(c echo.Context) string {
	username, _, _ := c.Request().BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Casbinauth) CheckPermission(c echo.Context) bool {
	user := c.Get("recid").(string)
	method := c.Request().Method
	path := c.Request().URL.Path
	return a.Enforcer.Enforce(user, path, method)
}

/*
Package casbin provides middleware to enable ACL, RBAC, ABAC authorization support.
Simple example:
	package main
	import (
		"github.com/casbin/casbin"
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo-contrib/casbin" casbin-mw
	)
	func main() {
		e := echo.New()
		// Mediate the access for every request
		e.Use(casbin-mw.Middleware(casbin.NewEnforcer("auth_model.conf", "auth_policy.csv")))
		e.Logger.Fatal(e.Start(":1323"))
	}
Advanced example:
	package main
	import (
		"github.com/casbin/casbin"
		"github.com/labstack/echo/v4"
		"github.com/labstack/echo-contrib/casbin" casbin-mw
	)
	func main() {
		ce := casbin.NewEnforcer("auth_model.conf", "")
		ce.AddRoleForUser("alice", "admin")
		ce.AddPolicy(...)
		e := echo.New()
		echo.Use(casbin-mw.Middleware(ce))
		e.Logger.Fatal(e.Start(":1323"))
	}
*/
