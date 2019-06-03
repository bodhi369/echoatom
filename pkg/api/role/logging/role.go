package role

import (
	"time"

	"github.com/bodhi369/echoatom/pkg/api/role"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// New creates new role logging service
func New(svc role.Service, logger schemago.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents role logging service
type LogService struct {
	role.Service
	logger schemago.Logger
}

const name = "role"

// Query logging
func (ls *LogService) Query(c echo.Context, req *schemago.ReqRole) (resp []schemago.SRole, count int, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "query request", err,
			map[string]interface{}{
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Query(c, req)
}

// Get logging
func (ls *LogService) Get(c echo.Context, roleid string) (resp *schemago.RespRole, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "get request", err,
			map[string]interface{}{
				"roleId": roleid,
				"took":   time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Get(c, roleid)
}

// Create logging
func (ls *LogService) Create(c echo.Context, resprole *schemago.RespRole) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "create request", err,
			map[string]interface{}{
				"resprole": resprole,
				"took":     time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, resprole)
}

// Update logging
func (ls *LogService) Update(c echo.Context, resprole *schemago.RespRole) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "update request", err,
			map[string]interface{}{
				"resprole": resprole,
				"took":     time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, resprole)
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
