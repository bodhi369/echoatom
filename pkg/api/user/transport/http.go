package transport

import (
	"net/http"

	"github.com/bodhi369/echoatom/pkg/api/user"
	"github.com/bodhi369/echoatom/pkg/utl/schemago"

	"github.com/labstack/echo"
)

// HTTP represents user http service
type HTTP struct {
	svc user.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc user.Service, er *echo.Group) {
	h := HTTP{svc}
	// swagger:operation GET /v1/users user userList
	// ---
	// summary: list users
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
	//    "$ref": "#/responses/userResp"
	er.GET("/users", h.query)
	// swagger:operation GET /v1/users/{id} user userView
	// ---
	// summary: view user
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: userid
	//   type: string
	//   required: true
	// responses:
	//  200:
	//    "$ref": "#/responses/useridResp"
	er.GET("/users/:id", h.get)
	// swagger:operation POST /v1/users user userCreate
	// ---
	// summary: create user
	// description: Error Not Found (404) will be returned
	// responses:
	//  200:
	//    "$ref": "#/responses/usercreateResp"
	er.POST("/users", h.create)
	// swagger:operation POST /v1/users/{id} user userUpdated
	// ---
	// summary: update user
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: userid
	//   type: string
	//   required: true
	// - name: Body
	//   in: body
	//   schema:
	//     "$ref": "#/definitions/ReqUpdateUser"
	// responses:
	//  200:
	//    "$ref": "#/responses/usercreateResp"
	er.PUT("/users/:id", h.update)
	// swagger:operation DELETE /v1/users/{id} user userDelete
	// ---
	// summary: delete user
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: userid
	//   type: string
	//   required: true
	er.DELETE("/users/:id", h.delete)
	// swagger:operation PATCH /v1/users/{id}/enable user userEnable
	// ---
	// summary: enable user
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: userid
	//   type: string
	//   required: true
	er.PATCH("/users/:id/enable", h.enable)
	// swagger:operation PATCH /v1/users/{id}/disable user userDisable
	// ---
	// summary: disable user
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: id
	//   in: path
	//   description: userid
	//   type: string
	//   required: true
	er.PATCH("/users/:id/disable", h.disable)

}

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

type queryResp struct {
	Resp  []schemago.RespUser
	Count int
	Page  int
}

func (h *HTTP) query(c echo.Context) error {
	p := new(schemago.ReqUser)
	if err := c.Bind(p); err != nil {
		return err
	}
	usr, count, err := h.svc.Query(c, p)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, queryResp{usr, count, p.Page})
}

func (h *HTTP) get(c echo.Context) error {
	req := c.Param("id")

	usr, err := h.svc.Get(c, req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) create(c echo.Context) error {
	req := new(schemago.ReqCreateUser)

	if err := c.Bind(req); err != nil {

		return err
	}

	if req.Password != req.PasswordConfirm {
		return ErrPasswordsNotMaching
	}

	usr, err := h.svc.Create(c, req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) update(c echo.Context) error {
	req := new(schemago.ReqUpdateUser)
	if err := c.Bind(req); err != nil {
		return err
	}

	usr, err := h.svc.Update(c, req)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *HTTP) enable(c echo.Context) error {
	id := c.Param("id")

	if err := h.svc.Enable(c, id, true); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *HTTP) disable(c echo.Context) error {
	id := c.Param("id")

	if err := h.svc.Disable(c, id, false); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
