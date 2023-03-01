package controllers

import (
	"encoding/json"
	"errors"
	"ridler_payments/models"

	// "ridler_payments/requests"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// AgentsController operations for Agents
type AgentsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AgentsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetBalance", c.GetBalance)
}

// Post ...
// @Title Post
// @Description create Agents
// @Param	body		body 	models.AgentDetailsRequest	true		"body for Agents content"
// @Success 201 {int} models.Agents
// @Failure 403 body is empty
// @router / [post]
func (c *AgentsController) Post() {
	var v models.AgentDetailsRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	c_by := v.UserId

	q := models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, Username: v.Msisdn, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	if _, err := models.AddAgents(&q); err == nil {
		c.Ctx.Output.SetStatus(201)
		agentNumber := "A" + strconv.Itoa(int(q.AgentId))

		q.AgentNumber = agentNumber

		if err := models.UpdateAgentsById(&q); err == nil {
			c.Data["json"] = q

			var balM = models.Balances{AgentId: int(q.AgentId), Balance: 0}

			if _, err := models.AddBalances(&balM); err == nil {
				c.Ctx.Output.SetStatus(201)
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = err.Error()
		}

	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Agents by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Agents
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AgentsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAgentsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Agents
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Agents
// @Failure 403
// @router / [get]
func (c *AgentsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAgents(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetBalance ...
// @Title Get Agent Balance
// @Description get Agent balance by id
// @Param	AgentId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Balances
// @Failure 403 :AgentId is empty
// @router /get-agent-balance/:AgentId [get]
func (c *AgentsController) GetBalance() {
	idStr := c.Ctx.Input.Param(":AgentId")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetBalanceByAgentId(int(id))

	logs.Info("Agent balance is ")
	logs.Info(v)

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Agents
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Agents	true		"body for Agents content"
// @Success 200 {object} models.Agents
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AgentsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Agents{AgentId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAgentsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Agents
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AgentsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteAgents(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
