package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type StatusCodes_20230109_191234 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &StatusCodes_20230109_191234{}
	m.Created = "20230109_191234"

	migration.Register("StatusCodes_20230109_191234", m)
}

// Run the migrations
func (m *StatusCodes_20230109_191234) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE status_codes(`status_code_id` int(11) NOT NULL AUTO_INCREMENT,`code` varchar(255) NOT NULL,`description` varchar(255) DEFAULT NULL,`message` varchar(255) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`status_code_id`))")
}

// Reverse the migrations
func (m *StatusCodes_20230109_191234) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `status_codes`")
}
