package controllers

import (
	"encoding/json"
	// "errors"
	"ridler_payments/models"
	// "strconv"
	// "strings"

	beego "github.com/beego/beego/v2/server/web"
)

// FundController operations for Requests
type FundController struct {
	beego.Controller
}

// URLMapping ...
func (c *FundController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create Balances
// @Param	body		body 	models.BalanceRequest	true		"body for Balances content"
// @Success 201 {int} models.Balances
// @Failure 403 body is empty
// @router / [post]
func (c *FundController) Post() {
	var v models.BalanceRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var z = models.Balances{AgentId: v.AgentId, Balance: v.Balance, BalanceId: 0}

	if err := models.UpdateBalancesByAgentId(&z, v.Action); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
