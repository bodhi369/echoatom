package gormplug

import (
	"github.com/bodhi369/echoatom/pkg/utl/config"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New creates new database connection to a postgres database
func New(dbCfg *config.Configuration) (*gorm.DB, error) {

	db, err := gorm.Open(dbCfg.DB.DriverName, dbCfg.DB.PSN)
	checkErr(err)
	// 全局禁用表名复数
	db.SingularTable(true)

	db.LogMode(true)
	return db, nil
}
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
