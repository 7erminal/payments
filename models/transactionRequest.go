package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type TransactionRequest struct {
	SendingAgentId  int64
	SendingBranchId int
	Amount          float32
	SenderName      string `orm:"size(255)"`
	SenderNumber    string `orm:"size(255)"`
	ReceiverName    string `orm:"size(255)"`
	ReceiverNumber  string `orm:"size(255)"`
	// ReceiverAddress string `orm:"size(255)"`
	TransactionType string `orm:"size(255)"`
	Source          string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(TransactionRequest))
}
