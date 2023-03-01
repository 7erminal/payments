package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CashOuts_20230222_185328 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CashOuts_20230222_185328{}
	m.Created = "20230222_185328"

	migration.Register("CashOuts_20230222_185328", m)
}

// Run the migrations
func (m *CashOuts_20230222_185328) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE cashOuts(`cashout_id` int(11) NOT NULL AUTO_INCREMENT,`transaction_id` int(11) NOT NULL,`receiving_branch_id` int(11) NOT NULL,`amount` int(11) NOT NULL,`status_code` int(11) NOT NULL,`receiver_name` varchar(100) NOT NULL,`receiver_number` varchar(100) NOT NULL,`receiver_id_number` varchar(100) DEFAULT NULL,`transaction_code` varchar(40) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT 0,PRIMARY KEY (`cashout_id`), FOREIGN KEY (transaction_id) REFERENCES transactions(transaction_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (receiving_branch_id) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CashOuts_20230222_185328) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `cashOuts`")
}
