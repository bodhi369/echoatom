package mssql

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/bodhi369/echoatom/pkg/utl/schemago"
)

// NewMenu returns a new Menu database instance
func NewMenu() *Menu {
	return &Menu{}
}

// Menu represents the client for Menu table
type Menu struct{}

// Query query Menus on database
func (u *Menu) Query(db *gorm.DB, reqmenu *schemago.ReqMenu) ([]schemago.RespMenu, int, error) {

	var menus = make([]schemago.SMenu, 0)
	pagination := reqmenu.PaginationReq.TransformMS()
	sqlHead := "select * from (select ROW_NUMBER() OVER (order by id) as rowid,* from s_menu"
	sqlWhere := " ([deleted_at] IS NULL) "
	if v := reqmenu.Code; v != "" {
		sqlWhere = sqlWhere + " and (code like %" + v + "%) "
	}
	if v := reqmenu.Name; v != "" {
		sqlWhere = sqlWhere + " and (name like %" + v + "%) "
	}
	if v := reqmenu.Router; v != "" {
		sqlWhere = sqlWhere + " and (router like %" + v + "%) "
	}
	if v := reqmenu.Parentid; v != "" {
		sqlWhere = sqlWhere + " and (parentid like %" + v + "%) "
	}
	if v := reqmenu.Hidden; v != -1 {
		sqlWhere = sqlWhere + " and (hidden = " + strconv.Itoa(v) + ")"
	}
	sqlEnd := ") TT where rowid between ? and ?"
	sqlText := sqlHead + " where " + sqlWhere + sqlEnd

	if err := db.Raw(sqlText, pagination.Limit, pagination.Offset).Scan(&menus).Error; err != nil {
		return nil, 0, err
	}
	var num int
	db.Model(&schemago.SMenu{}).Where(sqlWhere).Count(&num)
	resqMenu := make([]schemago.RespMenu, len(menus))
	for i, item := range menus {
		var resource schemago.SMenuResources
		db.Find(&resource, "menuid = ?", item.Recid)

		resqMenu[i] = schemago.RespMenu{Recid: item.Recid, Code: item.Code, Name: item.Name, Seq: item.Seq,
			Icon: item.Icon, Router: item.Router, Parentid: item.Parentid, Hidden: item.Hidden, Usercode: item.Usercode,
			Resources: resource.ToRespResource()}
	}
	return resqMenu, num, nil
}

// Get get roleMenu on database
func (u *Menu) Get(db *gorm.DB, menuid string) (*schemago.RespMenu, error) {
	var menu = new(schemago.SMenu)
	var menuresource = new(schemago.SMenuResources)
	if err := db.Where("recid = ?", menuid).Find(&menu).Error; err != nil {
		return nil, err
	}

	if err := db.Where("menuid = ?", menuid).Find(&menuresource).Error; err != nil {
		return nil, err
	}

	resp := &schemago.RespMenu{Recid: menu.Recid, Code: menu.Code, Name: menu.Name, Seq: menu.Seq,
		Icon: menu.Icon, Router: menu.Router, Parentid: menu.Parentid, Hidden: menu.Hidden, Usercode: menu.Usercode,
		Resources: menuresource.ToRespResource()}
	return resp, nil
}

// Create create
func (u *Menu) Create(db *gorm.DB, respMenu *schemago.RespMenu) error {
	menu := respMenu.ToSchemaMenu()
	menuResources := respMenu.ToSchemaMenuResources()
	tx := db.Begin()
	if err := tx.Create(&menu).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, item := range menuResources {
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// compareRoleMenu 对比并获取需要新增，修改，删除的菜单项
func (u *Menu) updateMenuResource(oldList, newList schemago.SMenuResources) (clist, dlist, ulist schemago.SMenuResources) {
	for _, nitem := range newList {
		exists := false
		for _, oitem := range oldList {
			if oitem.Code == nitem.Code {
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
			if nitem.Code == oitem.Code {
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

// updateRoleMenus 更新菜单数据
func (u *Menu) updateMenuResources(db *gorm.DB, oldList, newList schemago.SMenuResources) error {

	clist, dlist, ulist := u.updateMenuResource(oldList, newList)
	tx := db.Begin()
	for _, item := range clist {
		result := tx.Create(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range dlist {
		result := tx.Where("menuid=? AND code=?", item.Menuid, item.Code).Delete(schemago.SMenuResource{})
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, item := range ulist {
		result := tx.Model(&item).Where("menuid=? AND code=?", item.Menuid, item.Code).Omit("menuid", "code").Updates(&item)
		if err := result.Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// Update update
func (u *Menu) Update(db *gorm.DB, respMenu *schemago.RespMenu) error {
	menu := respMenu.ToSchemaMenu()
	newmenuresources := respMenu.ToSchemaMenuResources()

	oldmenuresources := make(schemago.SMenuResources, 0)
	if err := db.Find(&oldmenuresources, "menuid = ? ", menu.Recid).Error; err != nil {
		return err
	}
	if err := u.updateMenuResources(db, oldmenuresources, newmenuresources); err != nil {
		return err
	}
	return nil
}

// Delete delete Menu on database
func (u *Menu) Delete(db *gorm.DB, menuid string) error {

	tx := db.Begin()
	if err := tx.Where("menuid = ?", menuid).Delete(&schemago.SMenuResource{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("recid = ?", menuid).Delete(&schemago.SMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
