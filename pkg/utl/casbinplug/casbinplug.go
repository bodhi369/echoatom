package casbinplug

import (
	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/jinzhu/gorm"
)

// LoadRoleMenuPolicy 加载角色对应权限 p
func LoadRoleMenuPolicy(db *gorm.DB, sRoleMenu *schemago.SRoleMenu, model model.Model) error {
	var respRoleMenu = sRoleMenu.ToRespondRoleMenu()
	menuResource := make([]schemago.SMenuResource, 0)

	if err := db.Where("menuid = ? and code in (?)", respRoleMenu.Menuid, respRoleMenu.Resources).Find(&menuResource).Error; err != nil {
		return err
	}

	for _, item := range menuResource {
		var lineText = "p, " + sRoleMenu.Roleid + ", " + item.Path + ", " + item.Method
		persist.LoadPolicyLine(lineText, model)
	}
	return nil
}

// RemoveRoleMenuPolicy 清掉角色权限 p
func RemoveRoleMenuPolicy(roleid string, model model.Model) error {
	//清掉角色权限
	model.RemoveFilteredPolicy("p", "p", 0, roleid)
	return nil
}

// LoadUserRoleGroup 加载用户对应角色 g
func LoadUserRoleGroup(db *gorm.DB, userid string, model model.Model) error {
	userRoles := make([]schemago.SUserRole, 0)

	if err := db.Where("userid = ?", userid).Find(&userRoles).Error; err != nil {
		return err
	}
	for _, item := range userRoles {
		var lineText = "g, " + item.Userid + ", " + item.Roleid
		persist.LoadPolicyLine(lineText, model)
	}
	return nil

}

// RemoveUserRoleGroup 用户角色 g
func RemoveUserRoleGroup(userid string, model model.Model) error {
	//清掉角色权限
	model.RemoveFilteredPolicy("g", "g", 0, userid)
	return nil
}
