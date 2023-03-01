package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToBalanceIncTable_20230227_100315 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToBalanceIncTable_20230227_100315{}
	m.Created = "20230227_100315"

	migration.Register("AddColumnToBalanceIncTable_20230227_100315", m)
}

// Run the migrations
func (m *AddColumnToBalanceIncTable_20230227_100315) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE balance_inc ADD COLUMN `description` varChar(255) DEFAULT NULL AFTER `balance`")
}

// Reverse the migrations
func (m *AddColumnToBalanceIncTable_20230227_100315) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
