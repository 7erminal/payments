package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Credentials struct {
	Username string `orm:"size(20)"`
	Password string `orm:"size(20)"`
}

func init() {
	orm.RegisterModel(new(Credentials))
}
