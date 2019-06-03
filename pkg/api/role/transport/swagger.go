package transport

import "github.com/bodhi369/echoatom/pkg/utl/schemago"

// role response
// swagger:response roleResp
type swaggRoleResp struct {
	// in:body
	Body struct {
		*queryResp
	}
}

// role id response
// swagger:response roleidResp
type swaggRoleidResp struct {
	// in:body
	Body *schemago.RespRole
}

// role create
// swagger:parameters roleCreate
type swaggRolecReq struct {
	// in:body
	Body *schemago.RespRole
}

// role update
// swagger:parameters roleUpdate
type swaggRoleuReq struct {
	// in:body
	Body *schemago.RespRole
}
