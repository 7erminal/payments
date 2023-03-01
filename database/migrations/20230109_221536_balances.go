package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Balances_20230109_221536 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Balances_20230109_221536{}
	m.Created = "20230109_221536"

	migration.Register("Balances_20230109_221536", m)
}

// Run the migrations
func (m *Balances_20230109_221536) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE balances(`balance_id` int(11) NOT NULL AUTO_INCREMENT,`agent_id` int(11) DEFAULT NULL,`balance` float NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,PRIMARY KEY (`balance_id`), FOREIGN KEY (agent_id) REFERENCES agents(agent_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Balances_20230109_221536) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `balances`")
}
