// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/ubicaciones_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_lugar",
			beego.NSInclude(
				&controllers.TipoLugarController{},
			),
		),

		beego.NSNamespace("/ubicacion",
			beego.NSInclude(
				&controllers.UbicacionController{},
			),
		),

		beego.NSNamespace("/lugar_ubicacion",
			beego.NSInclude(
				&controllers.LugarUbicacionController{},
			),
		),

		beego.NSNamespace("/atributo_lugar",
			beego.NSInclude(
				&controllers.AtributoLugarController{},
			),
		),

		beego.NSNamespace("/valor_atributo_lugar",
			beego.NSInclude(
				&controllers.ValorAtributoLugarController{},
			),
		),

		beego.NSNamespace("/tipo_relacion_lugar",
			beego.NSInclude(
				&controllers.TipoRelacionLugarController{},
			),
		),

		beego.NSNamespace("/lugar",
			beego.NSInclude(
				&controllers.LugarController{},
			),
		),

		beego.NSNamespace("/relacion_lugares",
			beego.NSInclude(
				&controllers.RelacionLugaresController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
