package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCashoutsTable_20230227_100310 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCashoutsTable_20230227_100310{}
	m.Created = "20230227_100310"

	migration.Register("AddColumnToCashoutsTable_20230227_100310", m)
}

// Run the migrations
func (m *AddColumnToCashoutsTable_20230227_100310) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE cashOuts ADD COLUMN `receiving_agent_id` int(11) NOT NULL AFTER `transaction_id`, ADD FOREIGN KEY (receiving_agent_id) REFERENCES agents(agent_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToCashoutsTable_20230227_100310) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
