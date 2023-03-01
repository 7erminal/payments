package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Transactions struct {
	TransactionId     int64 `orm:"auto" orm: "omitempty"`
	RequestId         int64
	SendingAgentId    int64
	SendingBranchId   int
	Amount            float32
	SenderName        string `orm:"size(255)"`
	SenderNumber      string `orm:"size(255)"`
	StatusCode        int    `orm: "omitempty"`
	ReceiverName      string `orm:"size(255)"`
	ReceiverNumber    string `orm:"size(255)"`
	ReceivingAgentId  int64  `orm: "omitempty"`
	ReceivingBranchId int    `"omitempty"`
	TransactionCode   string `orm:"size(40)" orm: "omitempty"`
	Active            int    `orm: "omitempty"`
	TransactionType   int
	DateCreated       time.Time `orm:"type(datetime)" orm: "omitempty"`
	DateModified      time.Time `orm:"type(datetime)" orm: "omitempty"`
	CreatedBy         int       `orm: "omitempty"`
	ModifiedBy        int       `orm: "omitempty"`
}

func init() {
	orm.RegisterModel(new(Transactions))
}

// AddTransactions insert a new Transactions into database and returns
// last inserted Id on success.
func AddTransactions(m *Transactions) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTransactionsById retrieves Transactions by Id. Returns error if
// Id doesn't exist
func GetTransactionsById(id int64) (v *Transactions, err error) {
	o := orm.NewOrm()
	v = &Transactions{TransactionId: id}
	if err = o.QueryTable(new(Transactions)).Filter("TransactionId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetTransactionsById retrieves Transactions by Code. Returns error if
// Id doesn't exist
func GetTransactionsByCode(code string) (v *Transactions, err error) {
	o := orm.NewOrm()
	v = &Transactions{TransactionCode: code}
	if err = o.QueryTable(new(Transactions)).Filter("TransactionCode", code).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetTransactionsById retrieves Transactions by Agent Id. Returns error if
// Id doesn't exist
func GetTransactionsByAgentId(id int64) (v *Transactions, err error) {
	o := orm.NewOrm()
	v = &Transactions{TransactionId: id}
	if _, err = o.QueryTable(new(Transactions)).Filter("TransactionId", id).RelatedSel().All(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTransactions retrieves all Transactions matches certain condition. Returns empty list if
// no records exist
func GetAllTransactions(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Transactions))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Transactions
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTransactions updates Transactions by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransactionsById(m *Transactions) (err error) {
	o := orm.NewOrm()
	v := Transactions{TransactionId: m.TransactionId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTransactions deletes Transactions by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTransactions(id int64) (err error) {
	o := orm.NewOrm()
	v := Transactions{TransactionId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Transactions{TransactionId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
