package user

import (
	"time"

	"github.com/bodhi369/echoatom/pkg/api/user"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// New creates new user logging service
func New(svc user.Service, logger schemago.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents user logging service
type LogService struct {
	user.Service
	logger schemago.Logger
}

const name = "user"

// Query logging
func (ls *LogService) Query(c echo.Context, req *schemago.ReqUser) (resp []schemago.RespUser, count int, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Query user request", err,
			map[string]interface{}{
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Query(c, req)
}

// Get logging
func (ls *LogService) Get(c echo.Context, userid string) (resp *schemago.RespUser, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Get user request", err,
			map[string]interface{}{
				"req":  userid,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Get(c, userid)
}

// Create logging
func (ls *LogService) Create(c echo.Context, req *schemago.ReqCreateUser) (resp *schemago.SUser, err error) {
	defer func(begin time.Time) {
		req.Password = "xxx-redacted-xxx"
		ls.logger.Log(
			c,
			name, "Create user request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, userid string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete user request", err,
			map[string]interface{}{
				"req":  userid,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, userid)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *schemago.ReqUpdateUser) (resp *schemago.SUser, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update user request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}

// Enable Enable
func (ls *LogService) Enable(c echo.Context, userid string, active bool) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Enable user request", err,
			map[string]interface{}{
				"req":    userid,
				"active": active,
				"took":   time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Enable(c, userid, active)
}

// Disable Disable
func (ls *LogService) Disable(c echo.Context, userid string, active bool) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Disable user request", err,
			map[string]interface{}{
				"req":    userid,
				"active": active,
				"took":   time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Disable(c, userid, active)
}
