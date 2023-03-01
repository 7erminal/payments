package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToAgentsTable_20230223_165624 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToAgentsTable_20230223_165624{}
	m.Created = "20230223_165624"

	migration.Register("AddColumnToAgentsTable_20230223_165624", m)
}

// Run the migrations
func (m *AddColumnToAgentsTable_20230223_165624) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE agents ADD COLUMN `password` varchar(20) DEFAULT NULL AFTER `agent_name`")
}

// Reverse the migrations
func (m *AddColumnToAgentsTable_20230223_165624) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
