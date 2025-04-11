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
		beego.NSNamespace("/lugares",
			beego.NSInclude(
				&controllers.LugaresController{},
			),
		),
        beego.NSNamespace("/artistas",
			beego.NSInclude(
				&controllers.ArtistasController{},
			),
		),
		beego.NSNamespace("/administrador",
			beego.NSInclude(
				&controllers.AdministradorController{},
			),
		),
		beego.NSNamespace("/autor_coplas",
			beego.NSInclude(
				&controllers.Autor_coplasController{},
			),
		),
		beego.NSNamespace("/canciones",
			beego.NSInclude(
				&controllers.CancionesController{},
			),
		),
		beego.NSNamespace("/comida",
			beego.NSInclude(
				&controllers.ComidaController{},
			),
		),
		beego.NSNamespace("/coplas",
			beego.NSInclude(
				&controllers.CoplasController{},
			),
		),
		beego.NSNamespace("/credenciales",
			beego.NSInclude(
				&controllers.CredencialesController{},
			),
		),
		beego.NSNamespace("/cultura",
			beego.NSInclude(
				&controllers.CulturaController{},
			),
		),
		beego.NSNamespace("/estilo_musical",
			beego.NSInclude(
				&controllers.Estilo_musicalController{},
			),
		),
		beego.NSNamespace("/favoritos",
			beego.NSInclude(
				&controllers.FavoritosController{},
			),
		),
		beego.NSNamespace("/tipo_cultura",
			beego.NSInclude(
				&controllers.Tipo_culturaController{},
			),
		),
		beego.NSNamespace("/tipo_lugares",
			beego.NSInclude(
				&controllers.Tipo_lugaresController{},
			),
		),
		beego.NSNamespace("/trajes_tipicos",
			beego.NSInclude(
				&controllers.Trajes_tipicosController{},
			),
		),
		beego.NSNamespace("/usuario",
			beego.NSInclude(
				&controllers.UsuarioController{},
			),
		),


	)
	beego.AddNamespace(ns)
}
