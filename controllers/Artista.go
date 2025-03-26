package controllers

import (
	"github.com/sena_2824182/pija_music_mid/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"strconv"
)

// ArtistaController maneja las solicitudes relacionadas con artistas
type ArtistaController struct {
	beego.Controller
}

// Post maneja la creación de un nuevo artista
func (c *ArtistaController) Post() {
	var artista models.Artista
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &artista); err == nil {
		if id, err := models.AddArtista(&artista); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Artista creado correctamente", "ID": id}
		} else {
			c.Data["json"] = map[string]interface{}{"Success": false, "Message": err.Error()}
		}		
	} else {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error en los datos enviados"}
	}
	c.ServeJSON()
}

// Put maneja la actualización de un artista existente
func (c *ArtistaController) Put() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	artista, _ := models.GetArtista(id) 
	if artista != nil {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &artista); err == nil {
			if err := models.UpdateArtista(id, artista); err == nil {
				c.Data["json"] = map[string]interface{}{"Success": true, "Message": "Artista actualizado"}
			} else {
				c.Data["json"] = map[string]interface{}{"Success": false, "Message": err.Error()}
			}			
		} else {
			c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Error en los datos enviados"}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"Success": false, "Message": "Artista no encontrado"}
	}
	c.ServeJSON()
}