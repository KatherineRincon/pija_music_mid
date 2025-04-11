package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// Estilo_musicalController operations for Estilo_musical
type Estilo_musicalController struct {
	beego.Controller
}

// URLMapping ...
func (c *Estilo_musicalController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Estilo_musical
// @Param	body		body 	models.Estilo_musical	true		"body for Estilo_musical content"
// @Success 201 {object} models.Estilo_musical
// @Failure 403 body is empty
// @router / [post]
func (c *Estilo_musicalController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Estilo_musical by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Estilo_musical
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Estilo_musicalController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Estilo_musical
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Estilo_musical
// @Failure 403
// @router / [get]
func (c *Estilo_musicalController) GetAll() {
	fmt.Println("get all de estilo musical")

}

// Put ...
// @Title Put
// @Description update the Estilo_musical
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Estilo_musical	true		"body for Estilo_musical content"
// @Success 200 {object} models.Estilo_musical
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Estilo_musicalController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Estilo_musical
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Estilo_musicalController) Delete() {

}
