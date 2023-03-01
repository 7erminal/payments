package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Transactions_20230109_184909 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Transactions_20230109_184909{}
	m.Created = "20230109_184909"

	migration.Register("Transactions_20230109_184909", m)
}

// Run the migrations
func (m *Transactions_20230109_184909) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE transactions(`transaction_id` int(11) NOT NULL AUTO_INCREMENT,`request_id` int(11) DEFAULT NULL,`sending_branch_id` int(11) DEFAULT NULL,`amount` float NOT NULL,`sender_name` varchar(255) NOT NULL,`sender_number` varchar(255) NOT NULL,`status_code` int(11) DEFAULT NULL,`receiver_name` varchar(255) NOT NULL,`receiver_number` varchar(255) NOT NULL,`receiving_branch_id` int(11) DEFAULT NULL,`transaction_code` varchar(40) NOT NULL,`active` int(11) DEFAULT 0,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`transaction_id`), FOREIGN KEY (request_id) REFERENCES requests(request_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (sending_branch_id) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (receiving_branch_id) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Transactions_20230109_184909) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `transactions`")
}
