package main

import (
	"log"

	schemago "github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// mysql    github.com/go-sql-driver/mysql
	// mymysql          github.com/ziutek/mymysql/godrv
	// postgres         github.com/lib/pq
	// mssql    github.com/denisenkom/go-mssqldb
	// var dbType = "postgres"
	// var psn = `postgres://atomgo:atomgo123@192.168.1.14:5432/atom?sslmode=disable`
	// var dbType = "mssql"
	// var psn = `server=192.168.1.14;user id=sa;password=9999;database=casbin_init;encrypt=disable`
	//db, err := gorm.Open("mssql", "server=192.168.1.14;user id=sa;password=9999;database=casbin_init;encrypt=disable")
	db, err := gorm.Open("postgres", `postgres://atomgo:atomgo123@192.168.1.14:5432/atom?sslmode=disable`)
	defer db.Close()
	checkErr(err)
	// 全局禁用表名复数
	db.SingularTable(true)

	db.LogMode(true)
	//&schemago.SMenuAction{},
	db.AutoMigrate(&schemago.SUser{}, &schemago.SUserRole{}, &schemago.SMenu{}, &schemago.SMenuResource{}, &schemago.SRole{}, &schemago.SRoleMenu{})

	// 清空数据
	sql := `		
			truncate table s_user;
			truncate table s_user_role;
			truncate table s_menu;
			truncate table s_menu_resource;
			truncate table s_role;
			truncate table s_role_menu;
	       `
	err = db.Exec(sql).Error
	checkErr(err)
	// 初始数据
	menus := []schemago.SMenu{
		{Recid: "dda07027-f16d-4d80-a4fe-0f9acbbf75cb", Code: "S10", Name: "首页", Seq: "1000000", Icon: "dashboard", Router: "/dashboard", Hidden: false, Usercode: "admin"},
		{Recid: "82afda8c-f890-4ad2-acb7-4076d54032a3", Code: "S1100", Name: "系统管理", Seq: "1100000", Icon: "setting", Hidden: false, Usercode: "admin"},
		{Recid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "S1101", Name: "菜单管理", Seq: "1101000", Icon: "solution", Router: "/system/menu", Parentid: "82afda8c-f890-4ad2-acb7-4076d54032a3", Hidden: false, Usercode: "admin"},
		{Recid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "S1102", Name: "角色管理", Seq: "1102000", Icon: "audit", Router: "/system/role", Parentid: "82afda8c-f890-4ad2-acb7-4076d54032a3", Hidden: false, Usercode: "admin"},
		{Recid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "S1103", Name: "用户管理", Seq: "1103000", Icon: "user", Router: "/system/user", Parentid: "82afda8c-f890-4ad2-acb7-4076d54032a3", Hidden: false, Usercode: "admin"},
	}
	for _, item := range menus {
		db.Create(&item)
	}

	resources := []schemago.SMenuResource{
		{Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "query", Name: "查询菜单数据", Method: "GET", Path: "/v1/menus"},
		{Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "get", Name: "精确查询菜单数据", Method: "GET", Path: "/v1/menus/:id"},
		{Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "create", Name: "创建菜单数据", Method: "POST", Path: "/v1/menus"},
		{Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "update", Name: "更新菜单数据", Method: "PUT", Path: "/v1/menus/:id"},
		{Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Code: "delete", Name: "删除菜单数据", Method: "DELETE", Path: "/v1/menus/:id"},

		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "query", Name: "查询角色数据", Method: "GET", Path: "/v1/roles"},
		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "get", Name: "精确查询角色数据", Method: "GET", Path: "/v1/roles/:id"},
		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "create", Name: "创建角色数据", Method: "POST", Path: "/v1/roles"},
		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "update", Name: "更新角色数据", Method: "PUT", Path: "/v1/roles/:id"},
		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "delete", Name: "删除角色数据", Method: "DELETE", Path: "/v1/roles/:id"},
		{Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Code: "queryMenu", Name: "查询菜单数据", Method: "GET", Path: "/v1/menus"},

		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "query", Name: "查询用户数据", Method: "GET", Path: "/v1/users"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "get", Name: "精确查询用户数据", Method: "GET", Path: "/v1/users/:id"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "create", Name: "创建用户数据", Method: "POST", Path: "/v1/users"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "update", Name: "更新用户数据", Method: "PUT", Path: "/v1/users/:id"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "delete", Name: "删除用户数据", Method: "DELETE", Path: "/v1/users/:id"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "disable", Name: "禁用用户数据", Method: "PATCH", Path: "/v1/users/:id/disable"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "enable", Name: "启用用户数据", Method: "PATCH", Path: "/v1/users/:id/enable"},
		{Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Code: "queryRole", Name: "查询角色数据", Method: "GET", Path: "/v1/roles"},
	}
	for _, item := range resources {
		db.Create(&item)
	}

	roles := []schemago.SRole{
		{Recid: "6f246299-8ad1-43e1-b043-396416d8e3da", Code: "admin", Name: "管理员", Seq: "000001", Usercode: "admin"},
		{Recid: "0e556242-a85e-449d-8dde-80374d0a2446", Code: "test", Name: "测试", Seq: "000002", Usercode: "admin"},
	}
	for _, item := range roles {
		db.Create(&item)
	}

	roleMenus := []schemago.SRoleMenu{
		{Roleid: "6f246299-8ad1-43e1-b043-396416d8e3da", Menuid: "b59ec4b6-1624-4798-839d-7ffbde93ead7", Resource: "query,get,create,update,delete"},
		{Roleid: "6f246299-8ad1-43e1-b043-396416d8e3da", Menuid: "d5e74bf4-3a62-4e60-a4d5-f49424a18e84", Resource: "query,get,create,update,delete,queryMenu"},
		{Roleid: "6f246299-8ad1-43e1-b043-396416d8e3da", Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Resource: "query,get,create,update,delete,disable,enable,queryRole"},

		{Roleid: "0e556242-a85e-449d-8dde-80374d0a2446", Menuid: "08c958c1-f9c8-4e8c-80fc-b30913992ab0", Resource: "query,get,create,update,delete,disable,enable,queryRole"},
	}
	for _, item := range roleMenus {
		db.Create(&item)
	}

	users := []schemago.SUser{
		{Recid: "ebb44431-cfdf-4619-b3dd-6e99c47f7883", Code: "admin", Name: "admin", NickName: "管理员", Password: "$2a$10$UQrsfan0CzJkUsV9sUBw/uJH8/hTPyZNbzzTT78hBl/0hZ0yW6n0m", Active: true},
	}
	for _, item := range users {
		db.Create(&item)
	}

	userRole := []schemago.SUserRole{
		{Userid: "ebb44431-cfdf-4619-b3dd-6e99c47f7883", Roleid: "6f246299-8ad1-43e1-b043-396416d8e3da"},
	}
	for _, item := range userRole {
		db.Create(&item)
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
