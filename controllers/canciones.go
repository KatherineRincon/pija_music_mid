package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"
)

// CancionesController operations for Canciones
type CancionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CancionesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Canciones
// @Param	body		body 	models.Canciones	true		"body for Canciones content"
// @Success 201 {object} models.Canciones
// @Failure 403 body is empty
// @router / [post]
func (c *CancionesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Canciones by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Canciones
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CancionesController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Canciones
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Canciones
// @Failure 403
// @router / [get]
func (c *CancionesController) GetAll() {
	fmt.Println("get all de canciones")

	canciones_byte,_:= services.Metodo_get("host_api","Canciones?limit=0")
	JsonCanciones,_:= services.ProcesarJson(canciones_byte)
	fmt.Println("Json Canciones",JsonCanciones)

	SoloCanciones:= JsonCanciones["Data"]











	
	c.Data["json"] = map[string]interface{}{
			"Success": true,
			"Status":  200,
			"Message": "Consulta de Canciones",
			"Data":    JsonCanciones,
		}

		c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Canciones
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Canciones	true		"body for Canciones content"
// @Success 200 {object} models.Canciones
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CancionesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Canciones
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CancionesController) Delete() {

}
