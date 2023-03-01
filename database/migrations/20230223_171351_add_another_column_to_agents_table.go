package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddAnotherColumnToAgentsTable_20230223_171351 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddAnotherColumnToAgentsTable_20230223_171351{}
	m.Created = "20230223_171351"

	migration.Register("AddAnotherColumnToAgentsTable_20230223_171351", m)
}

// Run the migrations
func (m *AddAnotherColumnToAgentsTable_20230223_171351) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE agents ADD COLUMN `username` varchar(20) DEFAULT NULL AFTER `agent_name`")
}

// Reverse the migrations
func (m *AddAnotherColumnToAgentsTable_20230223_171351) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
