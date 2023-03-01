package controllers

import (
	"encoding/json"
	"errors"
	"ridler_payments/models"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/crypto/bcrypt"
)

// CredentialsController operations for Credentials
type CredentialsController struct {
	beego.Controller
}

// URLMapping ...
func (c *CredentialsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Validate", c.Validate)
}

// Post ...
// @Title Post
// @Description create Credentials
// @Param	body		body 	models.Credentials	true		"body for Credentials content"
// @Success 201 {int} models.Credentials
// @Failure 403 body is empty
// @router / [post]
func (c *CredentialsController) Post() {
	var v models.Credentials
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddCredentials(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Credentials by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Credentials
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CredentialsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCredentialsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// validate ...
// @Title validate
// @Description create Credentials
// @Param	body		body 	models.CredentialsRequest	true		"body for Credentials content"
// @Success 201 {int} models.Credentials
// @Failure 403 body is empty
// @router /validate [post]
func (c *CredentialsController) Validate() {
	var v models.CredentialsRequest
	// var q models.Agents
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if a, err := models.GetCredentialsByUsername(v.Username); err == nil {
		logs.Info("Username found")
		logs.Info(v.Username)
		logs.Info(a.Password)

		// Compare the stored hashed password, with the hashed version of the password that was received
		if err = bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(v.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			c.Data["json"] = err.Error()
		} else {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = a
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Credentials
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Credentials
// @Failure 403
// @router / [get]
func (c *CredentialsController) GetAll() {
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

	l, err := models.GetAllCredentials(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Credentials
// @Param	agentid		path 	string	true		"The id you want to update"
// @Param	body		body 	models.CredentialsRequest	true		"body for Credentials content"
// @Success 200 {object} models.Credentials
// @Failure 403 :id is not int
// @router /:agentid [put]
func (c *CredentialsController) Put() {
	idStr := c.Ctx.Input.Param(":agentid")
	logs.Debug("Received ID is ")
	logs.Info(idStr)
	id, _ := strconv.ParseInt(idStr, 0, 64)

	logs.Debug(id)
	v := models.Credentials{AgentId: id}
	// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Debug(v.Password)

	hashedPassword, errr := bcrypt.GenerateFromPassword([]byte(v.Password), 8)

	if errr == nil {
		logs.Debug(hashedPassword)

		v.Password = string(hashedPassword)

		logs.Debug("Sending", v.Password)

		// models.Agents{AgentName: v.AgentName, BranchId: v.BranchId, IdType: v.IdType, IdNumber: v.IdNumber, IsVerified: false, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: c_by, ModifiedBy: c_by}

	}

	if err := models.UpdateCredentialsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Credentials
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CredentialsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCredentials(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
