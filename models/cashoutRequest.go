package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type CashoutRequest struct {
	ReceivingAgentId  int64
	ReceivingBranchId int
	Amount            float32
	Code              string `orm:"size(255)"`
	ReceiverName      string `orm:"size(255)" orm: "omitempty"`
	ReceiverNumber    string `orm:"size(255)" orm: "omitempty"`
	TransactionType   string `orm:"size(255)" orm: "omitempty"`
	Source            string `orm:"size(255)" orm: "omitempty"`
}

func init() {
	orm.RegisterModel(new(CashoutRequest))
}
