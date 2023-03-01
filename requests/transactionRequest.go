package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type TransactionRequest struct {
	SendingBranchId   int
	Amount            float64
	SenderName        string `orm:"size(255)"`
	SenderNumber      string `orm:"size(255)"`
	ReceiverName      string `orm:"size(255)"`
	ReceiverNumber    string `orm:"size(255)"`
	ReceivingBranchId int
	Source string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(TransactionRequest))
}