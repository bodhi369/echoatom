package auth

import (
	sql "github.com/bodhi369/echoatom/pkg/api/auth/platform/mssql"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// New creates new iam service
func New(db *gorm.DB, ce *casbin.Enforcer, udb UserDB, j TokenGenerator, sec Securer) *Auth {
	return &Auth{
		db:  db,
		ce:  ce,
		udb: udb,
		tg:  j,
		sec: sec,
	}
}

// Initialize initializes auth application service
func Initialize(db *gorm.DB, ce *casbin.Enforcer, j TokenGenerator, sec Securer) *Auth {
	return New(db, ce, sql.NewUser(), j, sec)
}

// Service represents auth service interface
type Service interface {
	Authenticate(echo.Context, string, string) (*schemago.AuthToken, error)
	Refresh(echo.Context, string) (*schemago.RefreshToken, error)
	Logout(echo.Context, string) error
}

// Auth represents auth application service
type Auth struct {
	db  *gorm.DB
	ce  *casbin.Enforcer
	udb UserDB
	tg  TokenGenerator
	sec Securer
}

// UserDB represents user repository interface
type UserDB interface {
	FindByCode(*gorm.DB, string) (*schemago.SUser, error)
	FindByToken(*gorm.DB, string) (*schemago.SUser, error)
	Update(*gorm.DB, *schemago.SUser) error
}

// TokenGenerator represents token generator (jwt) interface
type TokenGenerator interface {
	GenerateToken(*schemago.SUser) (string, string, error)
}

// Securer represents security interface
type Securer interface {
	HashMatchesPassword(string, string) bool
	Token(string) string
}
