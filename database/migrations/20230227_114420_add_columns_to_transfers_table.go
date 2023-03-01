package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnsToTransfersTable_20230227_114420 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnsToTransfersTable_20230227_114420{}
	m.Created = "20230227_114420"

	migration.Register("AddColumnsToTransfersTable_20230227_114420", m)
}

// Run the migrations
func (m *AddColumnsToTransfersTable_20230227_114420) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE transfers ADD COLUMN `sending_balance_after` float DEFAULT 0 AFTER `sending_branch_id`, ADD COLUMN `sending_balance_before` float DEFAULT 0 AFTER `sending_branch_id`")
}

// Reverse the migrations
func (m *AddColumnsToTransfersTable_20230227_114420) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
