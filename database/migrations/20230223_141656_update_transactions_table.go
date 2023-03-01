package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UpdateTransactionsTable_20230223_141656 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UpdateTransactionsTable_20230223_141656{}
	m.Created = "20230223_141656"

	migration.Register("UpdateTransactionsTable_20230223_141656", m)
}

// Run the migrations
func (m *UpdateTransactionsTable_20230223_141656) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE transactions ADD COLUMN `transaction_type` int(11) NOT NULL AFTER `active`, ADD FOREIGN KEY (transaction_type) REFERENCES transaction_types(transaction_type_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *UpdateTransactionsTable_20230223_141656) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
