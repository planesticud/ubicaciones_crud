package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:AtributoLugarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:LugarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:RelacionLugaresController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoLugarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:TipoRelacionLugarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"] = append(beego.GlobalControllerRouter["github.com/udistrital/ubicaciones_crud/controllers:ValorAtributoLugarController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

}
