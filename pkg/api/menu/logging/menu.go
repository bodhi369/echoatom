package menu

import (
	"time"

	"github.com/bodhi369/echoatom/pkg/api/menu"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// New creates new menu logging service
func New(svc menu.Service, logger schemago.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents menu logging service
type LogService struct {
	menu.Service
	logger schemago.Logger
}

const name = "menu"

// Query logging
func (ls *LogService) Query(c echo.Context, reqmenu *schemago.ReqMenu) (resp []schemago.RespMenu, count int, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "query request", err,
			map[string]interface{}{
				"reqmenu": reqmenu,
				"took":    time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Query(c, reqmenu)
}

// Get logging
func (ls *LogService) Get(c echo.Context, menuid string) (resp *schemago.RespMenu, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "get request", err,
			map[string]interface{}{
				"menuid": menuid,
				"took":   time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Get(c, menuid)
}

// Create logging
func (ls *LogService) Create(c echo.Context, respMenu *schemago.RespMenu) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "create request", err,
			map[string]interface{}{
				"respMenu": respMenu,
				"took":     time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, respMenu)
}

// Update logging
func (ls *LogService) Update(c echo.Context, respMenu *schemago.RespMenu) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "update request", err,
			map[string]interface{}{
				"respMenu": respMenu,
				"took":     time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, respMenu)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, roleid string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "delete request", err,
			map[string]interface{}{
				"roleId": roleid,
				"took":   time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, roleid)
}
