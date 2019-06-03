package menu

import (
	sql "github.com/bodhi369/echoatom/pkg/api/menu/platform/mssql"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Service represents menu application interface
type Service interface {
	Query(echo.Context, *schemago.ReqMenu) ([]schemago.RespMenu, int, error)
	Get(echo.Context, string) (*schemago.RespMenu, error)
	Create(echo.Context, *schemago.RespMenu) error
	Update(echo.Context, *schemago.RespMenu) error
	Delete(echo.Context, string) error
}

// New creates new menu application service
func New(db *gorm.DB, udb UDB) *Menu {
	return &Menu{db: db, udb: udb}
}

// Initialize initalizes menu application service with defaults
func Initialize(db *gorm.DB) *Menu {
	return New(db, sql.NewMenu())
}

// Menu represents menu application service
type Menu struct {
	db  *gorm.DB
	udb UDB
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents menu repository interface
type UDB interface {
	Query(*gorm.DB, *schemago.ReqMenu) ([]schemago.RespMenu, int, error)
	Get(*gorm.DB, string) (*schemago.RespMenu, error)
	Create(*gorm.DB, *schemago.RespMenu) error
	Update(*gorm.DB, *schemago.RespMenu) error
	Delete(*gorm.DB, string) error
}
