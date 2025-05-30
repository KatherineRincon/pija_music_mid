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
	c.Mapping("Postloginad", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Postloginad ...
// @Title Create
// @Description create Administrador
// @Param	body		body 	models.Administrador	true		"body for Administrador content"
// @Success 201 {object} models.Administrador
// @Failure 403 body is empty
// @router /loginad [post]
func (c *AdministradorController) Postloginad() {
	var body_ingresa_admin map[string]interface{}
	var resultado_admin map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa_admin); err == nil {
		fmt.Println("json ingresa", body_ingresa_admin)
	}

	Cedula := fmt.Sprintf("%v", body_ingresa_admin["Cedula"])
	Contrasena := fmt.Sprintf("%v", body_ingresa_admin["Contraseña"])

	query := fmt.Sprintf("Cedula:%s,Contraseña:%s", Cedula, Contrasena)
	endpoint_admin := "Administrador?query=" + query

	body_admin_byte, _ := services.Metodo_get("hots_crud", endpoint_admin)
	body_admin_json, _ := services.ProcesarJson(body_admin_byte)

	data_admin := body_admin_json["Data"]
	fmt.Println("data", data_admin)

	if data_admin == nil {
		resultado_admin = map[string]interface{}{
			"Message": "Administrador no existente o credenciales incorrectas",
			"Caso":    1,
		}
	} else {
		resultado_admin = map[string]interface{}{
			"message": "Administrador y contraseña correctos",
			"Caso":    2,
			"Datos":   data_admin,
		}
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Consulta existosa",
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
