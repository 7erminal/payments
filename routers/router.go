// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ridler_payments/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),
		beego.NSNamespace("/agents",
			beego.NSInclude(
				&controllers.AgentsController{},
			),
		),
		beego.NSNamespace("/agent-credentials",
			beego.NSInclude(
				&controllers.CredentialsController{},
			),
		),
		beego.NSNamespace("/branches",
			beego.NSInclude(
				&controllers.BranchesController{},
			),
		),
		beego.NSNamespace("/transactions",
			beego.NSInclude(
				&controllers.TransactionsController{},
			),
		),
		beego.NSNamespace("/status-codes",
			beego.NSInclude(
				&controllers.Status_codesController{},
			),
		),
		beego.NSNamespace("/fund",
			beego.NSInclude(
				&controllers.FundController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
