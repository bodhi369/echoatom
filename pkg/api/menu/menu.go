package menu

import (
	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// Query provided by limit page code name ...
func (a *Menu) Query(c echo.Context, reqmenu *schemago.ReqMenu) ([]schemago.RespMenu, int, error) {
	u, count, err := a.udb.Query(a.db, reqmenu)
	if err != nil {
		return nil, 0, err
	}

	return u, count, nil
}

// Get provided by menuId
func (a *Menu) Get(c echo.Context, menuid string) (*schemago.RespMenu, error) {
	u, err := a.udb.Get(a.db, menuid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Create create menu
func (a *Menu) Create(c echo.Context, respMenu *schemago.RespMenu) error {
	return a.udb.Create(a.db, respMenu)
}

// Update update menu
func (a *Menu) Update(c echo.Context, respMenu *schemago.RespMenu) error {
	return a.udb.Update(a.db, respMenu)
}

// Delete provided by menuId
func (a *Menu) Delete(c echo.Context, menuid string) error {
	err := a.udb.Delete(a.db, menuid)
	if err != nil {
		return err
	}

	return nil
}
