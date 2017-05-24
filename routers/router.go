// @APIVersion 1.0.0
// @Title email-service API
// @Description email-service only serve email
// @Contact qsg@corex-tek.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"system-service/email-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/email",
			beego.NSInclude(
				&controllers.EmailController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
