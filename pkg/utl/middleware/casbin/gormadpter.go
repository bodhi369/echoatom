package casbinmw

import (
	"github.com/bodhi369/echoatom/pkg/utl/casbinplug"
	"github.com/bodhi369/echoatom/pkg/utl/schemago"
	"github.com/casbin/casbin/model"
	"github.com/jinzhu/gorm"
)

// Adapter represents the Gorm adapter for policy storage.
type Adapter struct {
	db *gorm.DB
}

// NewAdapter is the constructor for Adapter.
// dbSpecified is an optional bool parameter. The default value is false.
// It's up to whether you have specified an existing DB in dataSourceName.
// If dbSpecified == true, you need to make sure the DB in dataSourceName exists.
// If dbSpecified == false, the adapter will automatically create a DB named "casbin".
func NewAdapter(db *gorm.DB) *Adapter {
	return &Adapter{db: db}
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	var rolemenus schemago.SRoleMenus
	a.db.Find(&rolemenus)
	for _, item := range rolemenus {
		if err := casbinplug.LoadRoleMenuPolicy(a.db, &item, model); err != nil {
			return err
		}
	}
	var users []schemago.SUser
	a.db.Find(&users)
	for _, item := range users {
		if err := casbinplug.LoadUserRoleGroup(a.db, item.Recid, model); err != nil {
			return err
		}
	}

	return nil
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	// a.dropTable()
	// a.createTable()

	// for ptype, ast := range model["p"] {
	// 	for _, rule := range ast.Policy {
	// 		line := savePolicyLine(ptype, rule)
	// 		err := a.db.Create(&line).Error
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// for ptype, ast := range model["g"] {
	// 	for _, rule := range ast.Policy {
	// 		line := savePolicyLine(ptype, rule)
	// 		err := a.db.Create(&line).Error
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {

	return nil
}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {

	return nil
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {

	return nil
}
