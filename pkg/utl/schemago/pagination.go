package schemago

// Pagination constants
const (
	paginationDefaultLimit = 100
	paginationMaxLimit     = 500
)

// PaginationReq holds pagination http fields and tags
type PaginationReq struct {
	Limit int `query:"limit"`
	Page  int `query:"page" validate:"min=0"`
}

// Transform checks and converts http pagination into database pagination model
func (p *PaginationReq) Transform() *Pagination {
	if p.Limit < 1 {
		p.Limit = paginationDefaultLimit
	}

	if p.Limit > paginationMaxLimit {
		p.Limit = paginationMaxLimit
	}

	return &Pagination{Limit: p.Limit, Offset: p.Page * p.Limit}
}
func (p *PaginationReq) TransformMS() *Pagination {
	if p.Limit < 1 {
		p.Limit = paginationDefaultLimit
	}

	if p.Limit > paginationMaxLimit {
		p.Limit = paginationMaxLimit
	}

	return &Pagination{Limit: p.Page * p.Limit, Offset: (p.Page + 1) * p.Limit}
}

// Pagination holds paginations data
type Pagination struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}
