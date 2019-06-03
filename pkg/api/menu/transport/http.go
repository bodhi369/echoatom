package transport

import (
	"net/http"

	"github.com/bodhi369/echoatom/pkg/api/menu"
	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// HTTP represents menu http service
type HTTP struct {
	svc menu.Service
}

// NewHTTP creates new menu http service
func NewHTTP(svc menu.Service, er *echo.Group) {
	h := HTTP{svc}
	// swagger:operation GET /v1/menus menu menuList
	// ---
	// summary: list menus
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: limit
	//   in: query
	//   description: limit
	//   type: integer
	//   format: int32
	//   default: 100
	//   required: true
	// - name: page
	//   in: query
	//   description: page
	//   type: integer
	//   format: int32
	//   default: 1
	//   required: true
	// - name: code
	//   in: query
	//   description: code
	//   type: string
	//   required: false
	// - name: name
	//   in: query
	//   description: name
	//   type: string
	//   required: false
	// responses:
	//  200:
	//    "$ref": "#/responses/menuResp"
	er.GET("/menus", h.query)
	// swagger:operation GET /v1/menus/{id} menu menuView
	// ---
	// summary: view menu
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: menuid
	//   type: string
	//   required: true
	// responses:
	//  200:
	//    "$ref": "#/responses/menuidResp"
	er.GET("/menus/:id", h.get)
	// swagger:operation POST /v1/menus menu menuCreate
	// ---
	// summary: create menu
	// description: Error Not Found (404) will be returned
	// responses:
	//  200:
	//    "$ref": "#/responses/menuidResp"
	er.POST("/menus", h.create)
	// swagger:operation POST /v1/menus/{id} menu menuUpdated
	// ---
	// summary: update menu
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: menuid
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/RespMenu"
	// responses:
	//  200:
	//    "$ref": "#/responses/menuidResp"
	er.PUT("/menus/:id", h.update)
	// swagger:operation DELETE /v1/menus/{id} menu menuDelete
	// ---
	// summary: delete menu
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: menuid
	//   type: string
	//   required: true
	er.DELETE("/menus/:id", h.delete)

}

type queryResp struct {
	Resp  []schemago.RespMenu
	Count int
	Page  int
}

func (h *HTTP) query(c echo.Context) error {
	p := new(schemago.ReqMenu)
	if err := c.Bind(p); err != nil {
		return err
	}
	r, count, err := h.svc.Query(c, p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, queryResp{r, count, p.Page})
}

func (h *HTTP) get(c echo.Context) error {
	r, err := h.svc.Get(c, c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) create(c echo.Context) error {
	r := new(schemago.RespMenu)
	if err := c.Bind(r); err != nil {
		return err
	}
	return h.svc.Create(c, r)

}

func (h *HTTP) update(c echo.Context) error {
	r := new(schemago.RespMenu)
	if err := c.Bind(r); err != nil {
		return err
	}
	return h.svc.Update(c, r)
}

func (h *HTTP) delete(c echo.Context) error {
	return h.svc.Delete(c, c.Param("id"))
}
