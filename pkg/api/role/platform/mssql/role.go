package mssql

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"

	"github.com/bodhi369/echoatom/pkg/utl/casbinplug"
	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/labstack/echo"
)

// NewRole returns a new role database instance
func NewRole() *Role {
	return &Role{}
}

// Role represents the client for role table
type Role struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "role has used,can`t delete")
)

// Query query roles on database
func (u *Role) Query(db *gorm.DB, req *schemago.ReqRole) ([]schemago.SRole, int, error) {
	var roles = make([]schemago.SRole, 0)
	pagination := req.PaginationReq.TransformMS()
	sqlHead := "select * from (select ROW_NUMBER() OVER (order by id) as rowid,* from s_role"
	sqlWhere := " ([deleted_at] IS NULL) "
	if v := req.Code; v != "" {
		sqlWhere = sqlWhere + " and (code like %" + v + "%) "
	}
	if v := req.Name; v != "" {
		sqlWhere = sqlWhere + " and (name like %" + v + "%) "
	}
	sqlEnd := ") TT where rowid between ? and ?"
	sqlText := sqlHead + " where " + sqlWhere + sqlEnd

	if err := db.Raw(sqlText, pagination.Limit, pagination.Offset).Scan(&roles).Error; err != nil {
		return nil, 0, err
	}
	var num int
	db.Model(&schemago.SRole{}).Where(sqlWhere).Count(&num)
	return roles, num, nil
}

// Get get role on database
func (u *Role) Get(db *gorm.DB, roleid string) (*schemago.RespRole, error) {
	var role = new(schemago.SRole)
	var rolemenu = new(schemago.SRoleMenus)
	if err := db.Where("recid = ?", roleid).Find(&role).Error; err != nil {
		return nil, err
	}
	if err := db.Where("roleId = ?", roleid).Find(&rolemenu).Error; err != nil {
		return nil, err
	}
	resprole := role.ToRespondRole()
	resprole.Menus = rolemenu.ToRespondRoleMenus()
	return resprole, nil
}

// Create Create
func (u *Role) Create(db *gorm.DB, ce *casbin.Enforcer, resprole *schemago.RespRole) error {
	role := resprole.ToSchemaRole()
	rolemenu := resprole.ToSchemaRoleMenus()
	tx := db.Begin()
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, item := range rolemenu {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	//添加 policy 到 model 中
	casbinplug.RemoveRoleMenuPolicy(resprole.Recid, ce.GetModel())
	for _, item := range rolemenu {
		err := casbinplug.LoadRoleMenuPolicy(db, &item, ce.GetModel())
		if err != nil {
			return err
		}
	}
	// printPolicy(ce)
	return nil
}

// compareRoleMenu 对比并获取需要新增，修改，删除的菜单项
func (u *Role) compareRoleMenu(oldList, newList schemago.SRoleMenus) (clist, dlist, ulist schemago.SRoleMenus) {
	oldMap, newMap := oldList.ToMap(), newList.ToMap()

	for _, nitem := range newList {
		if _, ok := oldMap[nitem.Menuid]; ok {
			ulist = append(ulist, nitem)
			continue
		}
		clist = append(clist, nitem)
	}

	for _, oitem := range oldList {
		if _, ok := newMap[oitem.Menuid]; !ok {
			dlist = append(dlist, oitem)
		}
	}
	return
}

// updateRoleMenus 更新菜单数据
func (u *Role) updateRoleMenus(db *gorm.DB, oldrolemenus, newrolemenus schemago.SRoleMenus) error {

	clist, dlist, ulist := u.compareRoleMenu(oldrolemenus, newrolemenus)
	tx := db.Begin()
	for _, item := range clist {
		result := tx.Create(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range dlist {
		result := tx.Where("roleid=? AND menuid=?", item.Roleid, item.Menuid).Delete(schemago.SRoleMenu{})
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range ulist {
		result := tx.Model(&item).Where("roleid=? AND menuid=?", item.Roleid, item.Menuid).Omit("roleid", "menuid").Updates(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// Update update
func (u *Role) Update(db *gorm.DB, ce *casbin.Enforcer, resprole *schemago.RespRole) error {
	role := resprole.ToSchemaRole()
	newrolemenus := resprole.ToSchemaRoleMenus()

	oldrolemenus := make([]schemago.SRoleMenu, 0)
	if err := db.Find(&oldrolemenus, "roleid = ? ", role.Recid).Error; err != nil {
		return err
	}
	if err := u.updateRoleMenus(db, oldrolemenus, newrolemenus); err != nil {
		return err
	}
	//添加 policy 到 model 中
	casbinplug.RemoveRoleMenuPolicy(resprole.Recid, ce.GetModel())
	for _, item := range newrolemenus {
		err := casbinplug.LoadRoleMenuPolicy(db, &item, ce.GetModel())
		if err != nil {
			return err
		}
	}
	// printPolicy(ce)
	return nil
}

// Delete delete
func (u *Role) Delete(db *gorm.DB, ce *casbin.Enforcer, roleid string) error {
	var count = -1
	db.Model(&schemago.SUserRole{}).Where("roleid = ?", roleid).Count(&count)
	if count > 0 {
		return ErrAlreadyExists
	}
	tx := db.Begin()
	if err := tx.Where("roleid = ?", roleid).Delete(&schemago.SRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("recid = ?", roleid).Delete(&schemago.SRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	casbinplug.RemoveRoleMenuPolicy(roleid, ce.GetModel())
	// printPolicy(ce)
	return nil
}

// printPolicy 打印出当前 policy group
func printPolicy(ce *casbin.Enforcer) {
	policy := ce.GetPolicy()
	fmt.Println("p:", policy)
	groupPolicy := ce.GetGroupingPolicy()
	fmt.Println("g:", groupPolicy)
	allRoles := ce.GetAllRoles()
	fmt.Println("roles name:", allRoles)
	allSubjects := ce.GetAllSubjects()
	fmt.Println("sub:", allSubjects)
	allObjects := ce.GetAllObjects()
	fmt.Println("obj:", allObjects)
	allActions := ce.GetAllActions()
	fmt.Println("act:", allActions)
}
