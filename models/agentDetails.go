package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type AgentDetailsRequest struct {
	AgentName string `orm:"size(255)"`
	BranchId  int    `orm: "omitempty"`
	IdType    int    `orm: "omitempty"`
	IdNumber  string `orm:"size(50)" orm: "omitempty"`
	Source    string `orm:"size(50)"`
	Msisdn    string `orm:"size(50)"`
	UserId    int
}

func init() {
	orm.RegisterModel(new(AgentDetailsRequest))
}
