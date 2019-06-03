package schemago

import "github.com/jinzhu/gorm"

// SMenu 菜单
type SMenu struct {
	gorm.Model
	Recid    string `json:"recid" gorm:"type:varchar(40);not null;unique"`
	Code     string `json:"code" gorm:"type:varchar(50);not null;unique"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	Seq      string `json:"seq" gorm:"type:varchar(10)"`
	Icon     string `json:"icon" gorm:"type:varchar(20)"`
	Router   string `json:"router" gorm:"type:varchar(50)"`
	Parentid string `json:"parentid" gorm:"type:varchar(40)"`
	Hidden   bool   `json:"hidden"`
	Usercode string `json:"usercode" gorm:"type:varchar(40)"`
}

// SMenuAction 菜单
// type SMenuAction struct {
// 	gorm.Model
// 	Menuid string `json:"menuid" gorm:"type:varchar(40)"`
// 	Code   string `json:"code" gorm:"type:varchar(50)"`
// 	Name   string `json:"name" gorm:"type:varchar(50)"`
// }

// SMenuResource 菜单
type SMenuResource struct {
	gorm.Model
	Menuid string `json:"menuid" gorm:"type:varchar(40)"`
	Code   string `json:"code" gorm:"type:varchar(50)"`
	Name   string `json:"name" gorm:"type:varchar(50)"`
	Method string `json:"method" gorm:"type:varchar(50)"`
	Path   string `json:"path" gorm:"type:varchar(50)"`
}

//SMenuResources SMenuResources
type SMenuResources []*SMenuResource

// RespMenuResource RespMenuResource
type RespMenuResource struct {
	Code   string `json:"code" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Method string `json:"method" validate:"required"`
	Path   string `json:"path" validate:"required"`
}

//ToRespResource ToRespResource
func (s *SMenuResources) ToRespResource() RespMenuResources {
	resp := make(RespMenuResources, len(*s))
	for i, item := range *s {
		resp[i] = &RespMenuResource{
			Code:   item.Code,
			Name:   item.Name,
			Method: item.Method,
			Path:   item.Path,
		}
	}
	return resp
}

//RespMenuResources RespMenuResources
type RespMenuResources []*RespMenuResource

// RespMenu RespMenu
type RespMenu struct {
	Recid     string `json:"recid" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Seq       string `json:"seq" validate:"required"`
	Icon      string `json:"icon" validate:"required"`
	Router    string `json:"router" validate:"required"`
	Parentid  string `json:"parentid" validate:"omitempty"`
	Hidden    bool   `json:"hidden" validate:"omitempty"`
	Usercode  string `json:"usercode" validate:"required"`
	Resources RespMenuResources
}

// ToSchemaMenu ToSchemaMenu
func (m *RespMenu) ToSchemaMenu() *SMenu {
	return &SMenu{
		Recid:    m.Recid,
		Code:     m.Code,
		Name:     m.Name,
		Seq:      m.Seq,
		Icon:     m.Icon,
		Router:   m.Router,
		Parentid: m.Parentid,
		Hidden:   m.Hidden,
		Usercode: m.Usercode,
	}
}

// ToSchemaMenuResource ToSchemaMenuResource
func (m *RespMenuResource) ToSchemaMenuResource(menuid string) *SMenuResource {
	return &SMenuResource{
		Menuid: menuid,
		Code:   m.Code,
		Name:   m.Name,
		Method: m.Method,
		Path:   m.Path,
	}
}

// ToSchemaMenuResources ToSchemaMenuResources
func (m *RespMenu) ToSchemaMenuResources() SMenuResources {
	menuResources := make(SMenuResources, len(m.Resources))
	for i, item := range m.Resources {
		menuResources[i] = item.ToSchemaMenuResource(m.Recid)
	}
	return menuResources
}

// ReqMenu query 时的QueryParams
type ReqMenu struct {
	PaginationReq
	Code     string `json:"code" validate:"omitempty"`
	Name     string `json:"name" validate:"omitempty"`
	Router   string `json:"router" validate:"omitempty"`
	Parentid string `json:"parentid" validate:"omitempty"`
	Hidden   int    `json:"hidden" validate:"isdefault=-1,min=-1,max=1,omitempty"`
}
