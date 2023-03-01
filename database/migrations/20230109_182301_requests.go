package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Requests_20230109_182301 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Requests_20230109_182301{}
	m.Created = "20230109_182301"

	migration.Register("Requests_20230109_182301", m)
}

// Run the migrations
func (m *Requests_20230109_182301) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE requests(`request_id` int(11) NOT NULL AUTO_INCREMENT,`source` varchar(25) NOT NULL,`request` varchar(255) NOT NULL,`received_date` datetime DEFAULT CURRENT_TIMESTAMP,`response` varchar(255) DEFAULT NULL,`response_date` datetime NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime NOT NULL,`date_modified` datetime NOT NULL,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`request_id`))")
}

// Reverse the migrations
func (m *Requests_20230109_182301) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `requests`")
}
