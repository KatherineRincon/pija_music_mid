// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/sena_2824182/pija_music_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/Usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),
		beego.NSNamespace("/Artista",
		beego.NSInclude(
			&controllers.ArtistaController{},
		),
	),
		beego.NSNamespace("/Administrador",
			beego.NSInclude(
				&controllers.AdministradorController{},
			),
		),
		beego.NSNamespace("/Canciones",
			beego.NSInclude(
				&controllers.CancionesController{},
			),
		),
		beego.NSNamespace("/Comida",
			beego.NSInclude(
				&controllers.ComidaController{},
			),
		),
	)

	
		beego.NSNamespace("/Cultura",
			beego.NSInclude(
				&controllers.CulturaController{},
			),
		),


		beego.NSNamespace("/Trajes_tipicos",
			beego.NSInclude(
				&controllers.Trajes_tipicosController{},
			),
		),
	beego.AddNamespace(ns)
}
