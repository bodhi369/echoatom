package schemago

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// SRole 角色
type SRole struct {
	gorm.Model
	Recid    string `json:"recid" gorm:"type:varchar(40);not null;unique"`
	Code     string `json:"code" gorm:"type:varchar(50);not null;unique"`
	Name     string `json:"name" gorm:"type:varchar(50);not null;unique"`
	Seq      string `json:"seq" gorm:"type:varchar(10)"`
	Usercode string `json:"usercode" gorm:"type:varchar(40)"`
}

// SRoleMenu 角色对应的
type SRoleMenu struct {
	gorm.Model
	Roleid   string `json:"roleid" gorm:"type:varchar(40)"`
	Menuid   string `json:"menuid" gorm:"type:varchar(40)"`
	Resource string
	// Action   string
}

// SRoleMenus SRoleMenus
type SRoleMenus []SRoleMenu

// RespRole 返回
type RespRole struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Recid     string        `json:"recid" gorm:"type:varchar(40);not null;unique"`
	Code      string        `json:"code" gorm:"type:varchar(50);not null;unique"`
	Name      string        `json:"name" gorm:"type:varchar(50);not null;unique"`
	Seq       string        `json:"seq" gorm:"type:varchar(10)"`
	Usercode  string        `json:"usercode" gorm:"type:varchar(40)"`
	Menus     RespRoleMenus `json:"menus"`
}

// RespRoleMenu 返回
type RespRoleMenu struct {
	Menuid    string `json:"menuid" gorm:"type:varchar(40)"`
	Resources []string

	// Actions   []string
}

// RespRoleMenus 菜单
type RespRoleMenus []*RespRoleMenu

//ToRespondRole ToRespondRole
func (a *SRole) ToRespondRole() *RespRole {
	return &RespRole{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Recid:     a.Recid,
		Code:      a.Code,
		Name:      a.Name,
		Seq:       a.Seq,
		Usercode:  a.Usercode,
	}
}

// ToRespondRoleMenu 转换为角色菜单对象
func (a *SRoleMenu) ToRespondRoleMenu() *RespRoleMenu {
	item := &RespRoleMenu{
		Menuid: a.Menuid,
	}

	// if v := a.Action; &v != nil && v != "" {
	// 	item.Actions = strings.Split(v, ",")
	// }
	if v := a.Resource; &v != nil && v != "" {
		item.Resources = strings.Split(v, ",")
	}

	return item
}

// ToRespondRoleMenus 转换为角色菜单对象列表
func (a SRoleMenus) ToRespondRoleMenus() []*RespRoleMenu {
	list := make([]*RespRoleMenu, len(a))
	for i, item := range a {
		list[i] = item.ToRespondRoleMenu()
	}
	return list
}

// ToSchemaRole d
func (a *RespRole) ToSchemaRole() *SRole {
	return &SRole{
		Recid:    a.Recid,
		Code:     a.Code,
		Name:     a.Name,
		Seq:      a.Seq,
		Usercode: a.Usercode,
	}
}

// ToSchemaRoleMenu ToSchemaRoleMenu
func (a *RespRoleMenu) ToSchemaRoleMenu(roleID string) SRoleMenu {
	item := SRoleMenu{
		Roleid: roleID,
		Menuid: a.Menuid,
	}

	// var action string
	// if v := a.Actions; len(v) > 0 {
	// 	action = strings.Join(v, ",")
	// }
	// item.Action = action

	var resource string
	if v := a.Resources; len(v) > 0 {
		resource = strings.Join(v, ",")
	}
	item.Resource = resource

	return item
}

// ToSchemaRoleMenus dToSchemaRoleMenu
func (a RespRole) ToSchemaRoleMenus() []SRoleMenu {
	var rolemenu = make([]SRoleMenu, len(a.Menus))

	for i, item := range a.Menus {
		rolemenu[i] = item.ToSchemaRoleMenu(a.Recid)
	}
	return rolemenu
}

// ToMap 转换为键值映射
func (a SRoleMenus) ToMap() map[string]SRoleMenu {
	m := make(map[string]SRoleMenu)
	for _, item := range a {
		m[item.Menuid] = item
	}
	return m
}

// ReqRole  request role
type ReqRole struct {
	PaginationReq
	Code string `json:"code" validate:"omitempty"`
	Name string `json:"name" validate:"omitempty"`
}
