package transport

import (
	"net/http"

	"github.com/bodhi369/echoatom/pkg/api/role"
	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// HTTP represents role http service
type HTTP struct {
	svc role.Service
}

// NewHTTP creates role auth http service
func NewHTTP(svc role.Service, er *echo.Group) {
	h := HTTP{svc}
	// swagger:operation GET /v1/roles role roleList
	// ---
	// summary: list roles
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
	//    "$ref": "#/responses/roleResp"
	er.GET("/roles", h.query)
	// swagger:operation GET /v1/roles/{id} role roleView
	// ---
	// summary: view role
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: roleid
	//   type: string
	//   required: true
	// responses:
	//  200:
	//    "$ref": "#/responses/roleidResp"
	er.GET("/roles/:id", h.get)
	// swagger:operation POST /v1/roles role roleCreate
	// ---
	// summary: create role
	// description: Error Not Found (404) will be returned
	// responses:
	//  200:
	//    "$ref": "#/responses/roleidResp"
	er.POST("/roles", h.create)
	// swagger:operation POST /v1/roles/{id} role roleUpdated
	// ---
	// summary: update role
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: roleid
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/RespRole"
	// responses:
	//  200:
	//    "$ref": "#/responses/roleidResp"
	er.PUT("/roles/:id", h.update)
	// swagger:operation DELETE /v1/roles/{id} role roleDelete
	// ---
	// summary: delete role
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: roleid
	//   type: string
	//   required: true
	er.DELETE("/roles/:id", h.delete)

}

type queryResp struct {
	Resp  []schemago.SRole
	Count int
	Page  int
}

func (h *HTTP) query(c echo.Context) error {
	p := new(schemago.ReqRole)
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
	r := new(schemago.RespRole)
	if err := c.Bind(r); err != nil {
		return err
	}
	return h.svc.Create(c, r)

}

func (h *HTTP) update(c echo.Context) error {
	r := new(schemago.RespRole)
	if err := c.Bind(r); err != nil {
		return err
	}
	return h.svc.Update(c, r)
}

func (h *HTTP) delete(c echo.Context) error {
	return h.svc.Delete(c, c.Param("id"))
}
