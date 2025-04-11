package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:AdministradorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ArtistasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Autor_coplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Estilo_musicalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_culturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Tipo_lugaresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:Trajes_tipicosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/pija_music_mid/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
