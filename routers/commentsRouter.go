package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Activation_codesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:AgentsController"],
		beego.ControllerComments{
			Method:           "GetBalance",
			Router:           `/get-agent-balance/:AgentId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:BranchesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:agentid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:CredentialsController"],
		beego.ControllerComments{
			Method:           "Validate",
			Router:           `/validate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:FundController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:FundController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Id_typesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:RequestsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:Status_codesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "CashOut",
			Router:           `/cash-out`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetCashOutDetails",
			Router:           `/cash-out-details`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetAgentTransactionsWithAgentID",
			Router:           `/get-agent-transactions/:agentId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetAgentTransfersWithAgentID",
			Router:           `/get-agent-transfers/:agentId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "GetAgentCashoutsWithAgentID",
			Router:           `/get-agent-cashouts/:agentId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:TransactionsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"] = append(beego.GlobalControllerRouter["ridler_payments/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
