package role

import (
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// Query provided by limit page code name
func (a *Role) Query(c echo.Context, req *schemago.ReqRole) ([]schemago.SRole, int, error) {
	u, count, err := a.udb.Query(a.db, req)
	if err != nil {
		return nil, 0, err
	}

	return u, count, nil
}

// Get provided by roleId
func (a *Role) Get(c echo.Context, roleid string) (*schemago.RespRole, error) {
	u, err := a.udb.Get(a.db, roleid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Create create
func (a *Role) Create(c echo.Context, resprole *schemago.RespRole) error {
	return a.udb.Create(a.db, a.ce, resprole)
}

// Update update
func (a *Role) Update(c echo.Context, resprole *schemago.RespRole) error {
	return a.udb.Update(a.db, a.ce, resprole)
}

// Delete provided by roleId
func (a *Role) Delete(c echo.Context, roleid string) error {
	err := a.udb.Delete(a.db, a.ce, roleid)
	if err != nil {
		return err
	}
	return nil
}
