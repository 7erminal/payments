package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Branches_20230109_174944 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Branches_20230109_174944{}
	m.Created = "20230109_174944"

	migration.Register("Branches_20230109_174944", m)
}

// Run the migrations
func (m *Branches_20230109_174944) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE branches(`branch_id` int(11) NOT NULL AUTO_INCREMENT,`branch_name` varchar(255) NOT NULL,`branch_address` varchar(255) DEFAULT NULL,`branch_number` varchar(255) NOT NULL,`branch_description` varchar(255) DEFAULT NULL,`branch_location` varchar(255) DEFAULT NULL,`branch_phone_number` varchar(40) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`branch_id`))")
}

// Reverse the migrations
func (m *Branches_20230109_174944) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `branches`")
}
