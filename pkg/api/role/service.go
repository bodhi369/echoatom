package role

import (
	sql "github.com/bodhi369/echoatom/pkg/api/role/platform/mssql"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Service represents role application interface
type Service interface {
	Query(echo.Context, *schemago.ReqRole) ([]schemago.SRole, int, error)
	Get(echo.Context, string) (*schemago.RespRole, error)
	Create(echo.Context, *schemago.RespRole) error
	Update(echo.Context, *schemago.RespRole) error
	Delete(echo.Context, string) error
}

// New creates new role application service
func New(db *gorm.DB, ce *casbin.Enforcer, udb UDB, sec Securer) *Role {
	return &Role{db: db, ce: ce, udb: udb, sec: sec}
}

// Initialize initalizes role application service with defaults
func Initialize(db *gorm.DB, ce *casbin.Enforcer, sec Securer) *Role {
	return New(db, ce, sql.NewRole(), sec)
}

// Role represents role application service
type Role struct {
	db  *gorm.DB
	ce  *casbin.Enforcer
	udb UDB
	sec Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents role repository interface
type UDB interface {
	Query(*gorm.DB, *schemago.ReqRole) ([]schemago.SRole, int, error)
	Get(*gorm.DB, string) (*schemago.RespRole, error)
	Create(*gorm.DB, *casbin.Enforcer, *schemago.RespRole) error
	Update(*gorm.DB, *casbin.Enforcer, *schemago.RespRole) error
	Delete(*gorm.DB, *casbin.Enforcer, string) error
}
