package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"
)

// AdministradorController operations for Administrador
type AdministradorController struct {
	beego.Controller
}

// URLMapping ...
func (c *AdministradorController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Postlogin", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Postloginad ...
// @Title Login Administrador
// @Description Login de administrador con cédula y contraseña
// @Param	body		body 	map[string]interface{}	true	"Credenciales del administrador"
// @Success 200 {object} map[string]interface{}
// @Failure 400 el cuerpo está mal formado
// @Failure 500 error interno del servidor
// @router /login [post]
func (c *AdministradorController) Postlogin() {
	var body_ingresa_admin map[string]interface{}
	var resultado_admin map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa_admin); err != nil {
		c.CustomAbort(400, "Error al parsear el cuerpo de la solicitud")
		return
	}

	Cedula := fmt.Sprintf("%v", body_ingresa_admin["Cedula"])
	Contrasena := fmt.Sprintf("%v", body_ingresa_admin["Contraseña"])

	if Cedula == "" || Contrasena == "" {
		c.CustomAbort(400, "Cédula o contraseña vacía")
		return
	}

	query := fmt.Sprintf("Cedula:%s,Contraseña:%s", Cedula, Contrasena)
	endpoint_admin := "Administrador?query=" + query
	fmt.Println("jgkjj", endpoint_admin)

	body_admin_byte, err := services.Metodo_get("hots_crud", endpoint_admin)
	if err != nil {
		c.CustomAbort(500, "Error al consultar el servicio hots_crud")
		return
	}

	body_admin_json, err := services.ProcesarJson(body_admin_byte)
	if err != nil {
		c.CustomAbort(500, "Error al procesar la respuesta del servicio hots_crud")
		return
	}

	data_admin := body_admin_json["Data"]

	if data_admin == nil {
		resultado_admin = map[string]interface{}{
			"Message": "Administrador no existente o credenciales incorrectas",
			"Caso":    1,
		}
	} else {
		resultado_admin = map[string]interface{}{
			"Message": "Administrador y contraseña correctos",
			"Caso":    2,
			"Datos":   data_admin,
		}
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    resultado_admin,
	}
	c.ServeJSON()
}

// Post ...
// @Title Create
// @Description create Administrador
// @Param	body		body 	models.Administrador	true		"body for Administrador content"
// @Success 201 {object} models.Administrador
// @Failure 403 body is empty
// @router / [post]
func (c *AdministradorController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Administrador by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Administrador
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AdministradorController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Administrador
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Administrador
// @Failure 403
// @router / [get]
func (c *AdministradorController) GetAll() {
	fmt.Println("get all administrador")

}

// Put ...
// @Title Put
// @Description update the Administrador
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Administrador	true		"body for Administrador content"
// @Success 200 {object} models.Administrador
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AdministradorController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Administrador
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AdministradorController) Delete() {

}
