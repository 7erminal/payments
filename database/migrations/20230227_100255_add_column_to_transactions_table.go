package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToTransactionsTable_20230227_100255 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToTransactionsTable_20230227_100255{}
	m.Created = "20230227_100255"

	migration.Register("AddColumnToTransactionsTable_20230227_100255", m)
}

// Run the migrations
func (m *AddColumnToTransactionsTable_20230227_100255) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE transactions ADD COLUMN `sending_agent_id` int(11) NOT NULL AFTER `request_id`, ADD COLUMN `receiving_agent_id` int(11) DEFAULT NULL AFTER `receiver_number`, ADD FOREIGN KEY (sending_agent_id) REFERENCES agents(agent_id) ON UPDATE CASCADE ON DELETE NO ACTION, ADD FOREIGN KEY (receiving_agent_id) REFERENCES agents(agent_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToTransactionsTable_20230227_100255) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
