package user

import (
	sql "github.com/bodhi369/echoatom/pkg/api/user/platform/mssql"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, *schemago.ReqCreateUser) (*schemago.SUser, error)
	Query(echo.Context, *schemago.ReqUser) ([]schemago.RespUser, int, error)
	Get(echo.Context, string) (*schemago.RespUser, error)
	Delete(echo.Context, string) error
	Update(echo.Context, *schemago.ReqUpdateUser) (*schemago.SUser, error)
	Enable(echo.Context, string, bool) error
	Disable(echo.Context, string, bool) error
}

// New creates new user application service
func New(db *gorm.DB, ce *casbin.Enforcer, udb UDB, sec Securer) *User {
	return &User{db: db, ce: ce, udb: udb, sec: sec}
}

// Initialize initalizes user application service with defaults
func Initialize(db *gorm.DB, ce *casbin.Enforcer, sec Securer) *User {
	return New(db, ce, sql.NewUser(), sec)
}

// User represents user application service
type User struct {
	db  *gorm.DB
	ce  *casbin.Enforcer
	udb UDB
	sec Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(*gorm.DB, *casbin.Enforcer, *schemago.ReqCreateUser) (*schemago.SUser, error)
	Query(*gorm.DB, *schemago.ReqUser) ([]schemago.RespUser, int, error)
	Get(*gorm.DB, string) (*schemago.RespUser, error)
	Update(*gorm.DB, *casbin.Enforcer, *schemago.ReqUpdateUser) (*schemago.SUser, error)
	Delete(*gorm.DB, *casbin.Enforcer, string) error
	Enable(*gorm.DB, *casbin.Enforcer, string, bool) error
	Disable(*gorm.DB, *casbin.Enforcer, string, bool) error
}
