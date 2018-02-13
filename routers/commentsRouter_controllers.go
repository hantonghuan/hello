package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hello/controllers:MainController"] = append(beego.GlobalControllerRouter["hello/controllers:MainController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hello/controllers:MainController"] = append(beego.GlobalControllerRouter["hello/controllers:MainController"],
		beego.ControllerComments{
			Method: "Wedding",
			Router: `/wedding`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
