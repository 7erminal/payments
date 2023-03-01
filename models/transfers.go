package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Transfers struct {
	TransferId             int64 `orm:"auto"`
	TransactionId          int
	SendingAgentId         int64
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
	orm.RegisterModel(new(Transfers))
}

// AddTransfers insert a new Transfers into database and returns
// last inserted Id on success.
func AddTransfers(m *Transfers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTransfersById retrieves Transfers by Id. Returns error if
// Id doesn't exist
func GetTransfersById(id int64) (v *Transfers, err error) {
	o := orm.NewOrm()
	v = &Transfers{TransferId: id}
	if err = o.QueryTable(new(Transfers)).Filter("TransferId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetTransfersById retrieves Transfers by Id. Returns error if
// Id doesn't exist
func GetTransfersByCode(code string) (v *Transfers, err error) {
	o := orm.NewOrm()
	v = &Transfers{TransactionCode: code}

	if err = o.QueryTable(new(Transfers)).Filter("TransactionCode", code).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTransfers retrieves all Transfers matches certain condition. Returns empty list if
// no records exist
func GetAllTransfers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Transfers))
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

	var l []Transfers
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

// UpdateTransfers updates Transfers by Id and returns error if
// the record to be updated doesn't exist
func UpdateTransfersById(m *Transfers) (err error) {
	o := orm.NewOrm()
	v := Transfers{TransferId: m.TransferId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTransfers deletes Transfers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTransfers(id int64) (err error) {
	o := orm.NewOrm()
	v := Transfers{TransferId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Transfers{TransferId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
