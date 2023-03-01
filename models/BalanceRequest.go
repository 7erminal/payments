package models

import (
	"github.com/beego/beego/v2/client/orm"
	// "strconv"
)

type BalanceRequest struct {
	BalanceId int64 `orm:"auto" orm: "omitempty"`
	AgentId   int
	Balance   float32 `orm: "omitempty"`
	Action    string
}

func init() {
	orm.RegisterModel(new(BalanceRequest))
}
