package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Transfers_20230222_184619 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Transfers_20230222_184619{}
	m.Created = "20230222_184619"

	migration.Register("Transfers_20230222_184619", m)
}

// Run the migrations
func (m *Transfers_20230222_184619) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE transfers(`transfer_id` int(11) NOT NULL AUTO_INCREMENT,`transaction_id` int(11) NOT NULL,`sending_branch_id` int(11) NOT NULL,`amount` int(11) NOT NULL,`sender_name` varchar(100) NOT NULL,`sender_number` varchar(50) NOT NULL,`status_code` int(11) NOT NULL,`receiver_name` varchar(100) NOT NULL,`receiver_number` varchar(100) NOT NULL,`transaction_code` varchar(40) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT 0,PRIMARY KEY (`transfer_id`), FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (sending_branch_id) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Transfers_20230222_184619) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `transfers`")
}
