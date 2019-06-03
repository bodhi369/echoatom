package transport

import "github.com/bodhi369/echoatom/pkg/utl/schemago"

// menu response
// swagger:response menuResp
type swaggMenuResp struct {
	// in:body
	Body struct {
		*queryResp
	}
}

// menu id response
// swagger:response menuidResp
type swaggMenuidResp struct {
	// in:body
	Body *schemago.RespMenu
}

// menu create
// swagger:parameters menuCreate
type swaggMenucReq struct {
	// in:body
	Body *schemago.RespMenu
}

// menu update
// swagger:parameters menuUpdate
type swaggMenuuReq struct {
	// in:body
	Body *schemago.RespMenu
}
