package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/sena_2824182/pija_music_mid/services"
)

// Controller para Autor_Coplas
type Autor_CoplasController struct {
	beego.Controller
}

// @Title CreateAutor_Coplas
// @Description create Autor_Coplas
// @Param	body		body 	models.Autor_Coplas	true		"body for Autor_Coplas content"
// @Success 200 {int} models.Autor_Coplas.Id
// @Failure 403 body is empty
// @router / [post]
func (u *Autor_CoplasController) Post() {
	var body_Autor_Coplas map[string]interface{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Autor_Coplas); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Procesar el Autor_Coplas
	json_Autor_Coplas_byte, err := json.Marshal(body_Autor_Coplas)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_Autor_Coplas, err := services.Metodo_post("services_post", json_Autor_Coplas_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error en el servicio POST"}
		u.ServeJSON()
		return
	}
	fmt.Println("Respuesta del POST:", string(response_Autor_Coplas))

	var Autor_Cloplas_response map[string]interface{}
	if err := json.Unmarshal(response_Autor_Coplas, &Autor_Cloplas_response); err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar la respuesta"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = Autor_Cloplas_response
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Autor_Coplas
// @Success 200 {object} models.Autor_Coplas
// @router / [get]
func (u *Autor_CoplasController) GetAll() {
	users := models.GetAllAutor_Coplas() // Corregido: ahora solo recibe un valor

	if len(users) == 0 { // Verifica si la lista está vacía
		u.Data["json"] = map[string]string{"error": "No hay Autor_Coplas registrados"}
	} else {
		u.Data["json"] = users
	}

	u.ServeJSON()
}

// @Title Get
// @Description get Autor_Coplas by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Autor_Coplas
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *Autor_CoplasController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.GetAllAutor_Coplas(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al obtener usuario"}
		} else {
			u.Data["json"] = user
		}
	} else {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the Autor_Coplas
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Autor_Coplas	true		"body for Autor_Coplas content"
// @Success 200 {object} models.Autor_Coplas
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *Autor_CoplasController) Put() {
	id_Autor_Coplas := u.Ctx.Input.Param(":uid")
	var body_Autor_Coplas map[string]interface{}

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Autor_Coplas); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Obtener el Autor_Coplas actual
	get_Autor_Coplas, err := services.Metodo_get("servicio", id_Autor_Coplas)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al obtener autor_coplas"}
		u.ServeJSON()
		return
	}
	fmt.Println("Autor_Coplas obtenido:", string(get_Autor_Coplas))

	json_Autor_Coplas, err := services.ProcesadorJson(get_Autor_Coplas)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar autor_coplas"}
		u.ServeJSON()
		return
	}
	fmt.Println("JSON del Autor_Cloplas:", json_Autor_Coplas)

	err = models.UpdateUsuario(id_Autor_Coplas, AutorActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar autor_coplas"}
		u.ServeJSON()
		return
	}

	// Actualizar Autor_Coplas
	json_byte, err := json.Marshal(Autor_CloplasActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_put, err := services.Metodo_put("servicio", id_Autor_Coplas, json_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar Autor_Coplas"}
	} else {
		u.Data["json"] = string(response_put)
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the Autor_Coplas
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *Autor_CoplasController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	if uid == "" {
		u.Data["json"] = map[string]string{"error": "Autor_Coplas ID is required"}
	} else {
		err := models.DeleteAutorCopla(:uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al eliminar Autor_Coplas"}
		} else {
			u.Data["json"] = "delete success!"
		}
	}
	u.ServeJSON()
}
