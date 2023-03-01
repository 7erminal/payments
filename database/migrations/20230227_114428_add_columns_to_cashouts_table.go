package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToCashoutsTable_20230227_114428 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToCashoutsTable_20230227_114428{}
	m.Created = "20230227_114428"

	migration.Register("AddColumnsToCashoutsTable_20230227_114428", m)
}

// Run the migrations
func (m *AddColumnsToCashoutsTable_20230227_114428) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE cashOuts ADD COLUMN `receiving_balance_after` float DEFAULT 0 AFTER `receiving_branch_id`, ADD COLUMN `receiving_balance_before` float DEFAULT 0 AFTER `receiving_branch_id`")
}

// Reverse the migrations
func (m *AddColumnsToCashoutsTable_20230227_114428) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
