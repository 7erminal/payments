package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type TransactionTypes_20230223_141543 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &TransactionTypes_20230223_141543{}
	m.Created = "20230223_141543"

	migration.Register("TransactionTypes_20230223_141543", m)
}

// Run the migrations
func (m *TransactionTypes_20230223_141543) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE transaction_types(`transaction_type_id` int(11) NOT NULL AUTO_INCREMENT,`transaction_type` varchar(50) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`transaction_type_id`))")
}

// Reverse the migrations
func (m *TransactionTypes_20230223_141543) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `transaction_types`")
}
