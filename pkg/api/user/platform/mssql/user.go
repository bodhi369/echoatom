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

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "user code already exists.")
)

// Query creates a new user on database
func (u *User) Query(db *gorm.DB, req *schemago.ReqUser) ([]schemago.RespUser, int, error) {
	users := make([]schemago.SUser, 0)
	pagination := req.PaginationReq.TransformMS()
	sqlHead := "select * from (select ROW_NUMBER() OVER (order by id) as rowid,* from s_user"
	sqlWhere := " ([deleted_at] IS NULL) "
	if v := req.Code; v != "" {
		sqlWhere = sqlWhere + " and (code like %" + v + "%) "
	}
	if v := req.Name; v != "" {
		sqlWhere = sqlWhere + " and (name like %" + v + "%) "
	}
	sqlEnd := ") TT where rowid between ? and ?"
	sqlText := sqlHead + " where " + sqlWhere + sqlEnd

	if err := db.Raw(sqlText, pagination.Limit, pagination.Offset).Scan(&users).Error; err != nil {
		return nil, 0, err
	}
	var num int
	db.Model(&schemago.SUser{}).Where(sqlWhere).Count(&num)

	respUsers := make([]schemago.RespUser, len(users))
	for i, item := range users {
		var userRoles schemago.SUserRoles
		if err := db.Where("userid = ?", item.Recid).Find(&userRoles).Error; err != nil {
			return nil, 0, err
		}
		respUsers[i] = schemago.RespUser{SUser: item, SUserRoles: userRoles}
	}
	return respUsers, num, nil
}

// Get creates a new user on database
func (u *User) Get(db *gorm.DB, userid string) (*schemago.RespUser, error) {
	users := schemago.SUser{}
	if err := db.Find(&users, "recid = ?", userid).Error; err != nil {
		return nil, err
	}

	var userRoles schemago.SUserRoles
	if err := db.Where("userid = ?", userid).Find(&userRoles).Error; err != nil {
		return nil, err
	}

	return &schemago.RespUser{SUser: users, SUserRoles: userRoles}, nil
}

// Create creates a new user on database
func (u *User) Create(db *gorm.DB, ce *casbin.Enforcer, req *schemago.ReqCreateUser) (*schemago.SUser, error) {
	var user = new(schemago.SUser)
	if !db.Where("code = ?", req.Code).Find(&user).RecordNotFound() {
		return nil, ErrAlreadyExists
	}
	user = req.ToSchemaSUser()
	userRoles := req.ToSchemaSUserRoles()

	tx := db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, item := range userRoles {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	casbinplug.RemoveUserRoleGroup(req.Recid, ce.GetModel())
	casbinplug.LoadUserRoleGroup(db, req.Recid, ce.GetModel())
	// printPolicy(ce)
	return user, nil
}

// Update updates user's contact info
// compareUpdateRole 对比并获取需要新增，修改，删除的角色数据
func (u *User) compareUpdateRole(oldList, newList []*schemago.SUserRole) (clist, dlist, ulist []*schemago.SUserRole) {
	for _, nitem := range newList {
		exists := false
		for _, oitem := range oldList {
			if oitem.Roleid == nitem.Roleid {
				exists = true
				ulist = append(ulist, nitem)
				break
			}
		}
		if !exists {
			clist = append(clist, nitem)
		}
	}

	for _, oitem := range oldList {
		exists := false
		for _, nitem := range newList {
			if nitem.Roleid == oitem.Roleid {
				exists = true
				break
			}
		}
		if !exists {
			dlist = append(dlist, oitem)
		}
	}

	return
}

// updateUserRoles 更新菜单数据
func (u *User) updateUserRoles(db *gorm.DB, oldList, newList []*schemago.SUserRole) error {

	clist, dlist, ulist := u.compareUpdateRole(oldList, newList)
	tx := db.Begin()
	for _, item := range clist {
		result := tx.Create(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range dlist {
		result := tx.Where("userid=? AND roleid=?", item.Userid, item.Roleid).Delete(schemago.SUserRole{})
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range ulist {
		result := tx.Model(&item).Where("userid=? AND roleid=?", item.Userid, item.Roleid).Omit("userid", "roleid").Updates(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// Update Update
func (u *User) Update(db *gorm.DB, ce *casbin.Enforcer, req *schemago.ReqUpdateUser) (*schemago.SUser, error) {

	oldUserRoles := make([]*schemago.SUserRole, 0)
	if err := db.Where("userid = ?", req.Recid).Find(&oldUserRoles).Error; err != nil {
		return nil, err
	}

	user := req.ToSchemaSUser()
	newUserRoles := req.ToSchemaSUserRoles()

	if err := u.updateUserRoles(db, oldUserRoles, newUserRoles); err != nil {
		return nil, err
	}
	casbinplug.RemoveUserRoleGroup(req.Recid, ce.GetModel())
	casbinplug.LoadUserRoleGroup(db, req.Recid, ce.GetModel())
	// printPolicy(ce)
	return user, nil
}

// Delete sets deleted_at for a user
func (u *User) Delete(db *gorm.DB, ce *casbin.Enforcer, userid string) error {

	if err := db.Where("userid = ?", userid).Delete(schemago.SUserRole{}).Error; err != nil {
		return err
	}
	if err := db.Where("recid = ?", userid).Delete(schemago.SUser{}).Error; err != nil {
		return err
	}
	casbinplug.RemoveUserRoleGroup(userid, ce.GetModel())
	// printPolicy(ce)
	return nil
}

// Enable disable
func (u *User) Enable(db *gorm.DB, ce *casbin.Enforcer, userid string, active bool) error {
	if err := db.Model(&schemago.SUser{}).Where("recid = ?", userid).Update("active", active).Error; err != nil {
		return err
	}
	casbinplug.RemoveUserRoleGroup(userid, ce.GetModel())
	casbinplug.LoadUserRoleGroup(db, userid, ce.GetModel())
	// printPolicy(ce)
	return nil
}

// Disable disable
func (u *User) Disable(db *gorm.DB, ce *casbin.Enforcer, userid string, active bool) error {
	if err := db.Model(&schemago.SUser{}).Where("recid = ?", userid).Update("active", active).Error; err != nil {
		return err
	}
	casbinplug.RemoveUserRoleGroup(userid, ce.GetModel())
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
