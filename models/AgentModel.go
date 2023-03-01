package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type AgentModel struct {
	AgentId      int64     `orm:"auto" orm: "omitempty"`
	AgentName    string    `orm:"size(255)"`
	BranchId     int       `orm: "omitempty"`
	AgentNumber  string    `orm:"size(255)" orm: "omitempty"`
	IdType       int       `orm: "omitempty"`
	IdNumber     string    `orm:"size(50)"`
	IsVerified   bool      `orm: "omitempty"`
	Active       int       `orm: "omitempty"`
	DateCreated  time.Time `orm:"type(datetime)" orm: "omitempty"`
	DateModified time.Time `orm:"type(datetime)" orm: "omitempty"`
	CreatedBy    int       `orm: "omitempty"`
	ModifiedBy   int       `orm: "omitempty"`
	Username     string    `orm:"size(255)" orm: "omitempty"`
	Password     string    `orm:"size(255)" orm: "omitempty"`
}

func init() {
	orm.RegisterModel(new(AgentModel))
}
