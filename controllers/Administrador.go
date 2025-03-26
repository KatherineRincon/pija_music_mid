package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
)

// Operations about Administrador
type AdministradorController struct {
	beego.Controller
}

// @Title Create
// @Description create Administrador
// @Param	body		body 	models.Administrador	true		"The Administrador content"
// @Success 200 {string} models.Administrador.Id
// @Failure 403 body is empty
// @router / [post]
func (o *AdministradorController) Post() {
    var ob models.Administrador
    json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

    admin := models.Administrador(ob) // Crea el administrador y obtiene el objeto
    id := admin.ID                    // Extrae el ID correctamente

    o.Data["json"] = map[string]string{"Id": id.Hex()}
    o.ServeJSON()
}




// @Title Get
// @Description find Administrador by id
// @Param	Id		path 	string	true		"the id you want to get"
// @Success 200 {Administrador} models.Administrador
// @Failure 403 :Id is empty
// @router /:Id [get]
func (o *AdministradorController) Get() { 
		Id := o.Ctx.Input.Param(":Id")
		if Id != "" {
			ob, err := models.GetAdministrador(Id)  // Llama a la función para obtener un administrador por ID
			if err != nil {
				o.Data["json"] = err.Error()
			} else {
				o.Data["json"] = ob
			}
		}
		o.ServeJSON()
}
// @Title GetAll
// @Description get all Administrador
// @Success 200 {Administrador} models.Administrador
// @Failure 403 :AdministradorId is empty
// @router / [get]
func (o *AdministradorController) GetAllAdmins() {
    obs := models.GetAllAdmins()  // Llama a la función para obtener todos los administradores
    o.Data["json"] = obs
    o.ServeJSON()
}


// @Title Update
// @Description update the Administrador
// @Param	Id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Administrador	true		"The body"
// @Success 200 {Administrador} models.Administrador
// @Failure 403 :AdministradorId is empty
// @router /:Id [put]
func (o *AdministradorController) Put() {
    var ob models.Administrador
    Id := o.Ctx.Input.Param(":Id")

    if Id == "" {
        o.Data["json"] = "ID inválido"
        o.ServeJSON()
        return
    }

    adminID, err := strconv.ParseInt(Id, 10, 64)
    if err != nil {
        o.Data["json"] = "ID inválido"
        o.ServeJSON()
        return
    }

    if err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob); err != nil {
        o.Data["json"] = "Error al procesar el JSON"
        o.ServeJSON()
        return
    }

    // Convertimos adminID a string antes de pasarlo a models.Update
	if err := models.UpdateAdministrador(strconv.FormatInt(adminID, 10), ob); err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "¡Actualización exitosa!"
	}
	
    o.ServeJSON()
}



// @Title Delete
// @Description delete the Administrador
// @Param	Id		path 	string	true		"The Id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /:Id [delete]
func (o *AdministradorController) Delete() {
    Id := o.Ctx.Input.Param(":Id")
    models.DeleteAdministrador(Id)  // Llama a la función para eliminar un administrador por ID
    o.Data["json"] = "delete success!"
    o.ServeJSON()
}
