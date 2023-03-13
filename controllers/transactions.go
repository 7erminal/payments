package controllers

import (
	"encoding/json"
	"errors"
	"ridler_payments/models"

	// "ridler_payments/requests"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// TransactionsController operations for Transactions
type TransactionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *TransactionsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAgentTransactionsWithAgentID", c.GetAgentTransactionsWithAgentID)
	c.Mapping("CashOut", c.CashOut)
	c.Mapping("GetCashOutDetails", c.GetCashOutDetails)
	c.Mapping("GetAgentTransfersWithAgentID", c.GetAgentTransfersWithAgentID)
	c.Mapping("GetAgentCashoutsWithAgentID", c.GetAgentCashoutsWithAgentID)
}

func RandStringBytes(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Post ...
// @Title Post
// @Description create Transactions
// @Param	body		body 	models.TransactionRequest	true		"body for Transactions content"
// @Success 201 {int} models.Transactions
// @Failure 403 body is empty
// @router / [post]
func (c *TransactionsController) Post() {
	var v models.TransactionRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var r = models.Requests{Source: v.Source, Request: string(c.Ctx.Input.RequestBody[:]), Active: 1, ReceivedDate: time.Now(), DateCreated: time.Now()}

	if _, err := models.AddRequests(&r); err == nil {
		txnCode := RandStringBytes(6) + "-" + strconv.Itoa(int(r.RequestId))

		transactionType := v.TransactionType

		if transactionType == "TRANSFER" {
			transactionType = "1"
		} else if transactionType == "CASHOUT" {
			transactionType = "2"
		} else if transactionType == "DEPOSIT" {
			transactionType = "3"
		}

		transactionTypeConv, _ := strconv.Atoi(transactionType)

		statusCode := 2002

		var t = models.Transactions{RequestId: r.RequestId, SendingAgentId: v.SendingAgentId, SendingBranchId: v.SendingBranchId, Amount: v.Amount, SenderName: v.SenderName, SenderNumber: v.SenderNumber, ReceiverName: v.ReceiverName, ReceiverNumber: v.ReceiverNumber, TransactionType: transactionTypeConv, Active: 1, DateCreated: time.Now(), CreatedBy: v.SendingBranchId, TransactionCode: txnCode, StatusCode: statusCode}
		if _, err := models.AddTransactions(&t); err == nil {
			logs.Info("No error inserting into transactions table")
			logs.Info("transaction type is ", transactionType)
			logs.Info("Transaction ID is ", t.TransactionId)
			if transactionType == "1" {
				// Check the balance of the sender
				senderBalance, s_err := models.GetBalanceByAgentId(int(v.SendingAgentId))

				logs.Info("Balance is ", senderBalance.Balance)

				if s_err != nil {
					c.Data["json"] = s_err.Error()
				} else {
					// If Sender has enough balance to transact
					if float32(senderBalance.Balance) > v.Amount {

						agent_, err := models.GetAgentsById(v.SendingAgentId)

						if err != nil {
							c.Data["json"] = s_err.Error()
						} else {
							// Saving in transfers
							var tr = models.Transfers{TransactionId: int(t.TransactionId), SendingAgent: agent_, SendingBranchId: v.SendingBranchId, Sending_balance_before: float32(senderBalance.Balance), Sending_balance_after: 0.0, Amount: v.Amount, SenderName: v.SenderName, SenderNumber: v.SenderNumber, ReceiverName: v.ReceiverName, ReceiverNumber: v.ReceiverNumber, TransactionCode: txnCode, Active: 2, DateCreated: time.Now(), CreatedBy: v.SendingBranchId, StatusCode: statusCode}

							// If save is successful
							if _, err := models.AddTransfers(&tr); err == nil {
								senderAfterBalance := models.Balances{AgentId: int(v.SendingAgentId), Balance: v.Amount}
								// senderBalance := models.Balances{AgentId: int(v.SendingAgentId), Balance: v.Amount}
								// Update sender's balance
								if err := models.UpdateBalancesByAgentId(&senderAfterBalance, "subtract"); err == nil {
									// var balUpdate = models.Transfers{TransferId: tr.TransferId}

									h_balance, err := models.GetBalanceIncById(2)

									if err != nil {
										c.Data["json"] = err.Error()
									} else {
										tr.Sending_balance_after = senderAfterBalance.Balance

										h_balance.Balance = h_balance.Balance + v.Amount

										if err := models.UpdateBalance_incById(h_balance); err == nil {
											// Update balance after in transfers table after sender balance has been updated
											if err := models.UpdateTransfersById(&tr); err == nil {
												c.Data["json"] = "OK"
												c.Ctx.Output.SetStatus(201)
												c.Data["json"] = t
											} else {
												c.Data["json"] = err.Error()
											}
										}
									}

								} else {
									c.Data["json"] = err.Error()
								}
							} else {
								// return errors.New("Error: Invalid order. Must be either [asc|desc]")
							}
						}
					}
				}
			} else if transactionType == "2" {
				transactionType = "2"
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else if transactionType == "3" {
				transactionType = "3"
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			}

		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// GetCashOutDetails ...
// @Title GetCashOutDetails
// @Description cashout Transaction details
// @Param	body		body 	models.CashoutRequest	true		"body for Transactions content"
// @Success 201 {int} models.Transactions
// @Success 200 {int} models.Transfers
// @Failure 403 body is empty
// @router /cash-out-details [post]
func (c *TransactionsController) GetCashOutDetails() {
	var v models.CashoutRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var r = models.Requests{Source: v.Source, Request: string(c.Ctx.Input.RequestBody[:]), Active: 1, ReceivedDate: time.Now(), DateCreated: time.Now()}

	if _, err := models.AddRequests(&r); err == nil {
		getTransaction, err := models.GetTransfersByCode(v.Code)

		logs.Info(getTransaction)

		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			if getTransaction.Active == 2 {
				c.Data["json"] = getTransaction
			} else {
				c.Data["json"] = "This transaction has already been withdrawn"
			}
		}
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// CashOut ...
// @Title CashOut
// @Description cashout Transactions
// @Param	body		body 	models.CashoutRequest	true		"body for Transactions content"
// @Success 201 {int} models.Transfers
// @Failure 403 body is empty
// @router /cash-out [post]
func (c *TransactionsController) CashOut() {
	var v models.CashoutRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	var r = models.Requests{Source: v.Source, Request: string(c.Ctx.Input.RequestBody[:]), Active: 1, ReceivedDate: time.Now(), DateCreated: time.Now()}

	if _, err := models.AddRequests(&r); err == nil {

		transactionType := v.TransactionType

		if transactionType == "TRANSFER" {
			transactionType = "1"
		} else if transactionType == "CASHOUT" {
			transactionType = "2"
		} else if transactionType == "DEPOSIT" {
			transactionType = "3"
		}

		logs.Info("About to go get code")

		statusCode := 2002

		getTransaction, err := models.GetTransfersByCode(v.Code)

		logs.Info(err)

		if getTransaction.Active == 2 {

			if err != nil {
				c.Data["json"] = err.Error()
			} else {
				logs.Info(getTransaction.Amount)
				logs.Info("vs", v.Amount)

				timeCreated := time.Now()

				agent_, err := models.GetAgentsById(v.ReceivingAgentId)

				if err != nil {
					c.Data["json"] = err.Error()
				} else {
					var cashout_ = models.CashOuts{TransactionId: getTransaction.TransactionId, ReceivingAgent: agent_, ReceivingBranchId: v.ReceivingBranchId, Receiving_balance_before: 0.0, Receiving_balance_after: 0.0, Amount: int(v.Amount), StatusCode: statusCode, ReceiverName: v.ReceiverName, ReceiverNumber: v.ReceiverNumber, TransactionCode: v.Code, DateCreated: timeCreated, CreatedBy: int(v.ReceivingAgentId), ModifiedBy: 0, Active: 2}

					// If save is successful
					if _, err := models.AddCashOuts(&cashout_); err == nil {
						if getTransaction.Amount == v.Amount {
							logs.Info("Amounts match")
							if getTransaction.ReceiverNumber == v.ReceiverNumber {
								// Check the balance of the receiver
								receiverBalance, r_err := models.GetBalanceByAgentId(int(v.ReceivingAgentId))

								if r_err != nil {
									c.Data["json"] = r_err.Error()
								} else {
									cashout_.Receiving_balance_before = receiverBalance.Balance

									receiverAfterBalance := models.Balances{AgentId: int(v.ReceivingAgentId), Balance: v.Amount}

									h_balance, err := models.GetBalanceIncById(2)

									if err != nil {
										c.Data["json"] = err.Error()
									} else {

										if err := models.UpdateBalancesByAgentId(&receiverAfterBalance, "add"); err == nil {
											// var balUpdate = models.Transfers{TransferId: tr.TransferId}
											h_balance.Balance = h_balance.Balance - v.Amount

											if err := models.UpdateBalance_incById(h_balance); err == nil {
												cashout_.Receiving_balance_after = receiverAfterBalance.Balance

												cashout_.DateModified = time.Now()
												cashout_.Active = 1
												cashout_.ModifiedBy = int(v.ReceivingAgentId)

												// Update balance after in transfers table after sender balance has been updated
												if err := models.UpdateCashOutsById(&cashout_); err == nil {
													getTransaction.Active = 1
													getTransaction.DateModified = time.Now()
													getTransaction.ModifiedBy = int(v.ReceivingAgentId)

													// Update transfers table to complete the transaction
													if err := models.UpdateTransfersById(getTransaction); err == nil {
														c.Data["json"] = "OK"
														c.Ctx.Output.SetStatus(201)
														c.Data["json"] = cashout_
													} else {
														c.Data["json"] = err.Error()
													}
												} else {
													c.Data["json"] = err.Error()
												}
											} else {
												c.Data["json"] = err.Error()
											}
										} else {
											c.Data["json"] = err.Error()
										}

									}
								}

							} else {
								c.Data["json"] = "Recipient number does not match"
							}
						} else {
							logs.Info("Amounts do not match")

							c.Data["json"] = "Amounts do not match"
						}
					} else {
						// return errors.New("Error: Invalid order. Must be either [asc|desc]")
						c.Data["json"] = err.Error()
					}
				}
			}
		} else {
			c.Data["json"] = "This transaction has already been withdrawn"
		}

	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Transactions by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Transactions
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TransactionsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTransactionsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAgentTransactionWithAgentID ...
// @Title Get Agent Transactions With AgentID
// @Description get Transactions by agent Id
// @Param	agentId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Transactions
// @Failure 403 :agentId is empty
// @router /get-agent-transactions/:agentId [get]
func (c *TransactionsController) GetAgentTransactionsWithAgentID() {
	idStr := c.Ctx.Input.Param(":agentId")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetTransactionsByAgentId(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAgentTransfersWithAgentID ...
// @Title Get Agent Transfers With AgentID
// @Description get Transfers by agent Id
// @Param	agentId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TransfersResponse
// @Failure 403 :agentId is empty
// @router /get-agent-transfers/:agentId [get]
func (c *TransactionsController) GetAgentTransfersWithAgentID() {
	idStr := c.Ctx.Input.Param(":agentId")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	v, err := models.GetTransfersByAgentId(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}

	c.ServeJSON()
}

// GetAgentCashoutsWithAgentID ...
// @Title Get Agent Cashouts With AgentID
// @Description get Cashouts by agent Id
// @Param	agentId		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.CashOuts
// @Failure 403 :agentId is empty
// @router /get-agent-cashouts/:agentId [get]
func (c *TransactionsController) GetAgentCashoutsWithAgentID() {
	idStr := c.Ctx.Input.Param(":agentId")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCashOutsByAgentId(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Transactions
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Transactions
// @Failure 403
// @router / [get]
func (c *TransactionsController) GetAll() {
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

	l, err := models.GetAllTransactions(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Transactions
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Transactions	true		"body for Transactions content"
// @Success 200 {object} models.Transactions
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TransactionsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Transactions{TransactionId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateTransactionsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Transactions
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TransactionsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteTransactions(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
