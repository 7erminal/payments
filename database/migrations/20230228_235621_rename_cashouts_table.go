package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RenameCashoutsTable_20230228_235621 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RenameCashoutsTable_20230228_235621{}
	m.Created = "20230228_235621"

	migration.Register("RenameCashoutsTable_20230228_235621", m)
}

// Run the migrations
func (m *RenameCashoutsTable_20230228_235621) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE `cashOuts` RENAME `cash_outs`")

}

// Reverse the migrations
func (m *RenameCashoutsTable_20230228_235621) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
