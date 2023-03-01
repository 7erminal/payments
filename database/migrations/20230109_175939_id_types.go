package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type IdTypes_20230109_175939 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &IdTypes_20230109_175939{}
	m.Created = "20230109_175939"

	migration.Register("IdTypes_20230109_175939", m)
}

// Run the migrations
func (m *IdTypes_20230109_175939) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE id_types(`id_type_id` int(11) NOT NULL AUTO_INCREMENT,`id_type_name` varchar(255) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,PRIMARY KEY (`id_type_id`))")
}

// Reverse the migrations
func (m *IdTypes_20230109_175939) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `id_types`")
}
