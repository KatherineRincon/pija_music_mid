package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
)

// ComidaController controla las operaciones de Comida
type ComidaController struct {
	beego.Controller
}

// @Title Create
// @Description Crear una nueva Comida
// @Param	body	body	models.Comida	true	"El contenido de la Comida"
// @Success 200 {string} models.Comida.IdComida
// @Failure 400 Body vacío o formato incorrecto
// @router / [post]
func (c *ComidaController) Post() {
	var ob models.Comida
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err != nil {
		c.Data["json"] = map[string]string{"error": "Formato JSON no válido"}
		c.ServeJSON()
		return
	}

	createdComida, err := models.CreateComida(ob)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"Id": createdComida.IdComida.Hex()}
	}
	c.ServeJSON()
}

// @Title Get
// @Description Obtener una Comida por ID
// @Param	Id	path	string	true	"El ID de la Comida"
// @Success 200 {object} models.Comida
// @Failure 400 ID vacío o inválido
// @router /:Id [get]
func (c *ComidaController) Get() {
	id := c.Ctx.Input.Param(":Id")
	if id == "" {
		c.Data["json"] = map[string]string{"error": "Se requiere ID"}
		c.ServeJSON()
		return
	}

	ob, err := models.GetComida(id)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = ob
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description Obtener todas las Comidas
// @Success 200 {array} models.Comida
// @router / [get]
func (c *ComidaController) GetAllComidas() {
	obs, err := models.GetAllComidas()
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = obs
	}
	c.ServeJSON()
}

// @Title Update
// @Description Actualizar una Comida
// @Param	Id	path	string	true	"El ID de la Comida a actualizar"
// @Param	body	body	models.Comida	true	"El cuerpo de la Comida"
// @Success 200 {string} mensaje de éxito
// @Failure 400 ID inválido o datos incorrectos
// @router /:Id [put]
func (c *ComidaController) Put() {
	id := c.Ctx.Input.Param(":Id")
	if id == "" {
		c.Data["json"] = map[string]string{"error": "Se requiere ID"}
		c.ServeJSON()
		return
	}

	var ob models.Comida
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &ob); err != nil {
		c.Data["json"] = map[string]string{"error": "Formato JSON no válido"}
		c.ServeJSON()
		return
	}

	if err := models.UpdateComida(id, ob); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Actualización exitosa"}
	}
	c.ServeJSON()
}

// @Title Delete
// @Description Eliminar una Comida
// @Param	Id	path	string	true	"El ID de la Comida a eliminar"
// @Success 200 {string} eliminación exitosa
// @Failure 400 ID vacío o inválido
// @router /:Id [delete]
func (c *ComidaController) Delete() {
	id := c.Ctx.Input.Param(":Id")
	if id == "" {
		c.Data["json"] = map[string]string{"error": "Se requiere ID"}
		c.ServeJSON()
		return
	}

	if err := models.DeleteComida(id); err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "Eliminación exitosa"}
	}
	c.ServeJSON()
}
