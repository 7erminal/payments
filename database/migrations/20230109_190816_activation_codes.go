package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ActivationCodes_20230109_190816 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ActivationCodes_20230109_190816{}
	m.Created = "20230109_190816"

	migration.Register("ActivationCodes_20230109_190816", m)
}

// Run the migrations
func (m *ActivationCodes_20230109_190816) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE activation_codes(`activation_code_id` int(11) NOT NULL AUTO_INCREMENT,`activation_code` varchar(10) NOT NULL,`processed` tinyint(1) DEFAULT 0,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`activation_code_id`))")
}

// Reverse the migrations
func (m *ActivationCodes_20230109_190816) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `activation_codes`")
}
