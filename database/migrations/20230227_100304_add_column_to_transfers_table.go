package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToTransfersTable_20230227_100304 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToTransfersTable_20230227_100304{}
	m.Created = "20230227_100304"

	migration.Register("AddColumnToTransfersTable_20230227_100304", m)
}

// Run the migrations
func (m *AddColumnToTransfersTable_20230227_100304) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE transfers ADD COLUMN `sending_agent_id` int(11) NOT NULL AFTER `transaction_id`, ADD FOREIGN KEY (sending_agent_id) REFERENCES agents(agent_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToTransfersTable_20230227_100304) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
