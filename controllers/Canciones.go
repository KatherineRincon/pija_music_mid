package controllers

import (
	"encoding/json"
	
	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
)

// CancionesController handles operations for Canciones
type CancionesController struct {
	beego.Controller
}

// @Title Create
// @Description create Canciones
// @Param	body	body	models.Canciones	true	"The Canciones content"
// @Success 200 {string} models.Canciones.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CancionesController) Post() {
	var ob models.Canciones
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	createdCancion, err := models.CreateCanciones(ob)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"Id": createdCancion.ID.Hex()}
	}
	c.ServeJSON()
}

// @Title Get
// @Description find Canciones by id
// @Param	Id	path	string	true	"The id you want to get"
// @Success 200 {object} models.Canciones
// @Failure 403 :Id is empty
// @router /:Id [get]
func (c *CancionesController) Get() {
	Id := c.Ctx.Input.Param(":Id")
	if Id == "" {
		c.Data["json"] = map[string]string{"error": "ID is required"}
		c.ServeJSON()
		return
	}

	ob, err := models.GetCanciones(Id)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = ob
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Canciones
// @Success 200 {array} models.Canciones
// @router / [get]
func (c *CancionesController) GetAllCanciones() {
	obs := models.GetAllCanciones()
	c.Data["json"] = obs
	c.ServeJSON()
}

// @Title Update
// @Description update the Canciones
// @Param	Id	path	string	true	"The id you want to update"
// @Param	body	body	models.Canciones	true	"The body"
// @Success 200 {string} success message
// @Failure 403 :Id is empty
// @router /:Id [put]
func (c *CancionesController) Put() {
	Id := c.Ctx.Input.Param(":Id")
	if Id == "" {
		c.Data["json"] = map[string]string{"error": "ID is required"}
		c.ServeJSON()
		return
	}

	var ob models.Canciones
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid JSON format"}
		c.ServeJSON()
		return
	}

	if err := models.UpdateCanciones(Id, ob); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Update successful"}
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Canciones
// @Param	Id	path	string	true	"The Id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /:Id [delete]
func (c *CancionesController) Delete() {
	Id := c.Ctx.Input.Param(":Id")
	if Id == "" {
		c.Data["json"] = map[string]string{"error": "ID is required"}
		c.ServeJSON()
		return
	}

	if err := models.DeleteCanciones(Id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Delete successful"}
	}
	c.ServeJSON()
}
