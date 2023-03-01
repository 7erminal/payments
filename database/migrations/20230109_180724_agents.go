package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Agents_20230109_180724 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Agents_20230109_180724{}
	m.Created = "20230109_180724"

	migration.Register("Agents_20230109_180724", m)
}

// Run the migrations
func (m *Agents_20230109_180724) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE agents(`agent_id` int(11) NOT NULL AUTO_INCREMENT,`agent_name` varchar(255) NOT NULL,`branch_id` int(11) DEFAULT NULL,`branch_number` varchar(255) NOT NULL,`id_type` int(11) DEFAULT NULL,`id_number` varchar(50) DEFAULT NULL,`is_verified` tinyint(1) DEFAULT 0,`active` int(11) DEFAULT 1,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`agent_id`), FOREIGN KEY (branch_id) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE RESTRICT)")
}

// Reverse the migrations
func (m *Agents_20230109_180724) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `agents`")
}
