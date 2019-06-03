package transport

import (
	"net/http"

	"github.com/bodhi369/echoatom/pkg/api/auth"

	"github.com/labstack/echo"
)

// HTTP represents auth http service
type HTTP struct {
	svc auth.Service
}

// NewHTTP creates new auth http service
func NewHTTP(svc auth.Service, e *echo.Echo, mw echo.MiddlewareFunc) {
	h := HTTP{svc}
	// swagger:operation POST /login auth login
	// ---
	// summary: user login
	// description: Error Not Found (404) will be returned
	// responses:
	//  200:
	//    "$ref": "#/responses/loginResp"
	e.POST("/login", h.login)
	// swagger:operation GET /refresh/{token} auth refresh
	// ---
	// summary: refresh user token
	// description: Error Not Found (404) will be returned
	// parameters:
	// - name: token
	//   in: path
	//   description: refresh token
	//   type: string
	//   required: true
	// responses:
	//  200:
	//    "$ref": "#/responses/refreshResp"
	e.GET("/refresh/:token", h.refresh)
	// swagger:operation POST /login/exit auth logout
	// ---
	// summary: user logout
	// description: logout
	e.POST("/login/exit", h.logout, mw)
}

type credentials struct {
	Code     string `json:"code" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *HTTP) login(c echo.Context) error {
	cred := new(credentials)
	if err := c.Bind(cred); err != nil {
		return err
	}
	r, err := h.svc.Authenticate(c, cred.Code, cred.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) refresh(c echo.Context) error {
	r, err := h.svc.Refresh(c, c.Param("token"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}

func (h *HTTP) logout(c echo.Context) error {
	err := h.svc.Logout(c, c.Get("code").(string))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
