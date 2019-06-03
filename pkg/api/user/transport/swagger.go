package transport

import "github.com/bodhi369/echoatom/pkg/utl/schemago"

// user response
// swagger:response userResp
type swaggUserResp struct {
	// in:body
	Body struct {
		*queryResp
	}
}

// user id response
// swagger:response useridResp
type swaggUseridResp struct {
	// in:body
	Body struct {
		*schemago.RespUser
	}
}

// user create
// swagger:parameters userCreate
type swaggUsercReq struct {
	// in:body
	Body *schemago.ReqCreateUser
}

// user create response
// swagger:response usercreateResp
type swaggUsercreateResp struct {
	// in:body
	Body struct {
		*schemago.SUser
	}
}

// user update
// swagger:parameters userUpdate
type swaggUseruReq struct {
	// in:body
	Body *schemago.ReqUpdateUser
}
