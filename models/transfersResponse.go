package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TransfersResponse struct {
	TransferId             int64 `orm:"auto"`
	TransactionId          int
	SendingAgentId         *Agents `orm:"rel(fk)"`
	SendingBranchId        int
	Sending_balance_before float32
	Sending_balance_after  float32
	Amount                 float32
	SenderName             string `orm:"size(100)"`
	SenderNumber           string `orm:"size(50)"`
	StatusCode             int
	ReceiverName           string    `orm:"size(100)"`
	ReceiverNumber         string    `orm:"size(100)"`
	TransactionCode        string    `orm:"size(40)"`
	DateCreated            time.Time `orm:"type(datetime)" orm: "omitempty"`
	DateModified           time.Time `orm:"type(datetime)" orm: "omitempty"`
	CreatedBy              int
	ModifiedBy             int
	Active                 int
}

func init() {
	orm.RegisterModel(new(TransfersResponse))
}
