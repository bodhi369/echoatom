package auth

import (
	"time"

	"github.com/bodhi369/echoatom/pkg/api/auth"
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// New creates new auth logging service
func New(svc auth.Service, logger schemago.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents auth logging service
type LogService struct {
	auth.Service
	logger schemago.Logger
}

const name = "auth"

// Authenticate logging
func (ls *LogService) Authenticate(c echo.Context, code, password string) (resp *schemago.AuthToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Authenticate request", err,
			map[string]interface{}{
				"req":  code,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Authenticate(c, code, password)
}

// Refresh logging
func (ls *LogService) Refresh(c echo.Context, req string) (resp *schemago.RefreshToken, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Refresh request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Refresh(c, req)
}

// Logout logging
func (ls *LogService) Logout(c echo.Context, req string) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Logout request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Logout(c, req)
}
