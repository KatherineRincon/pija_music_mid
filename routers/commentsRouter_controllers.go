package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Artista"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Artista"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Carta_contadorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Artista"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id/:id_2",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Artista"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:PaisesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Artista"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}