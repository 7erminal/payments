package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type BalanceInc_20230109_221544 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &BalanceInc_20230109_221544{}
	m.Created = "20230109_221544"

	migration.Register("BalanceInc_20230109_221544", m)
}

// Run the migrations
func (m *BalanceInc_20230109_221544) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE balance_inc(`balance_inc_id` int(11) NOT NULL AUTO_INCREMENT,`balance` float NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,PRIMARY KEY (`balance_inc_id`))")
}

// Reverse the migrations
func (m *BalanceInc_20230109_221544) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `balance_inc`")
}
