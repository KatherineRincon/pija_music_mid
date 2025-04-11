package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// Trajes_tipicosController operations for Trajes_tipicos
type Trajes_tipicosController struct {
	beego.Controller
}

// URLMapping ...
func (c *Trajes_tipicosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Trajes_tipicos
// @Param	body		body 	models.Trajes_tipicos	true		"body for Trajes_tipicos content"
// @Success 201 {object} models.Trajes_tipicos
// @Failure 403 body is empty
// @router / [post]
func (c *Trajes_tipicosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Trajes_tipicos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Trajes_tipicos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Trajes_tipicosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Trajes_tipicos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Trajes_tipicos
// @Failure 403
// @router / [get]
func (c *Trajes_tipicosController) GetAll() {
	fmt.Println("get all de trajes tipicos")

}

// Put ...
// @Title Put
// @Description update the Trajes_tipicos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Trajes_tipicos	true		"body for Trajes_tipicos content"
// @Success 200 {object} models.Trajes_tipicos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Trajes_tipicosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Trajes_tipicos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Trajes_tipicosController) Delete() {

}
