package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"

	// "time"
	"github.com/astaxie/beego/logs"
	// "strconv"
)

type Balances struct {
	BalanceId int64 `orm:"auto" orm: "omitempty"`
	AgentId   int
	Balance   float32 `orm: "omitempty"`
}

func init() {
	orm.RegisterModel(new(Balances))
}

func AddBalances(m *Balances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBalanceById retrieves Balances by Id. Returns error if
// Id doesn't exist
func GetBalanceById(id int64) (v *Balances, err error) {
	o := orm.NewOrm()
	v = &Balances{BalanceId: id}
	if err = o.QueryTable(new(Balances)).Filter("AgentId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetBalanceByAgentId retrieves Balances by Agent Id. Returns error if
// Id doesn't exist
func GetBalanceByAgentId(id int) (v *Balances, err error) {
	o := orm.NewOrm()
	v = &Balances{AgentId: id}
	if err = o.QueryTable(new(Balances)).Filter("AgentId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBalances retrieves all Balances matches certain condition. Returns empty list if
// no records exist
func GetAllBalances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Balances))
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

	var l []Balances
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

// UpdateBalancesById updates Balances by Id and returns error if
// the record to be updated doesn't exist
func UpdateBalancesById(m *Balances) (err error) {
	o := orm.NewOrm()
	v := Balances{BalanceId: m.BalanceId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UpdateBalancesByAgentId updates Balances by AgentId and returns error if
// the record to be updated doesn't exist
func UpdateBalancesByAgentId(m *Balances, action string) (err error) {
	o := orm.NewOrm()
	v := &Balances{AgentId: m.AgentId}

	received_balance := m.Balance

	if err = o.QueryTable(new(Balances)).Filter("AgentId", m.AgentId).RelatedSel().One(v); err == nil {
		logs.Info(v.Balance)
		// logs.Info(err)
		// return err
	}

	final_balance := v.Balance

	if action == "add" {
		final_balance = v.Balance + received_balance
		logs.Info("Addition")
	} else if action == "subtract" {
		final_balance = v.Balance - received_balance
	}

	logs.Info("Result after adding balance in DB and received balance is ")
	logs.Info(fmt.Sprintf("%v", v.Balance) + " + " + fmt.Sprintf("%v", received_balance))
	logs.Info(final_balance)

	m.BalanceId = v.BalanceId
	m.Balance = final_balance

	logs.Info("New balance to be updated is ")
	logs.Info(m.Balance)
	logs.Info(m.AgentId)
	logs.Info(m.BalanceId)
	logs.Info(v.Balance)
	logs.Info(final_balance)

	// logs.Info(m.BalanceId)
	// ascertain id exists in the database
	if err = o.Read(v); err == nil {
		logs.Info("It exists")
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBalances deletes Balances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBalances(id int64) (err error) {
	o := orm.NewOrm()
	v := Balances{BalanceId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Balances{BalanceId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
