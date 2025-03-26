package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/sena_2824182/pija_music_mid/services"
)

// Controller para usuarios
type CulturaController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *CulturaController) Post() {
	var body_Danza map[string]interface{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Danza); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Procesar el usuario
	json_Danza_byte, err := json.Marshal(body_Danza)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_Danza, err := services.Metodo_post("services_post", json_Danza_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error en el servicio POST"}
		u.ServeJSON()
		return
	}
	fmt.Println("Respuesta del POST:", string(response_Danza))

	var Danza_response map[string]interface{}
	if err := json.Unmarshal(response_Danza, &Danza_response); err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar la respuesta"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = Danza_response
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *CulturaController) GetAll() {
	users := models.GetAllCultura() // Corregido: ahora solo recibe un valor

	if len(users) == 0 { // Verifica si la lista está vacía
		u.Data["json"] = map[string]string{"error": "No hay cultura registrados"}
	} else {
		u.Data["json"] = users
	}

	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *CulturaController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.GetCultura(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al obtener cultura"}
		} else {
			u.Data["json"] = user
		}
	} else {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *CulturaController) Put() {
	id_Cultura := u.Ctx.Input.Param(":uid")
	var body_Cultura map[string]interface{}

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Cultura); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Obtener el usuario actual
	get_Cultura, err := services.Metodo_get("servicio", id_Cultura)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al obtener cultura"}
		u.ServeJSON()
		return
	}
	fmt.Println("Cultura obtenido:", string(get_Cultura))

	json_Cultura, err := services.ProcesadorJson(get_Cultura)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar cultura"}
		u.ServeJSON()
		return
	}
	fmt.Println("JSON del cultura:", json_Cultura)

	// Construir nuevo JSON
	CulturaActualizado := models.Cultura{
		ID:       id_Cultura,
		Nombre_Cultura:   fmt.Sprintf("%v", body_Cultura["nombre_cultuta"]),
		Descripcion_Cultura:   fmt.Sprintf("%v", body_Cultura["descripcion_cultura"]),
		Imagen_video: fmt.Sprintf("%v", body_Cultura["imagen_video"]), // Asegúrate de validar esto
	}

	err = models.UpdateCultura(id_Cultura, CulturaActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar cultura"}
		u.ServeJSON()
		return
	}

	// Actualizar culrura
	json_byte, err := json.Marshal(CulturaActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_put, err := services.Metodo_put("servicio", id_Cultura, json_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar cultura"}
	} else {
		u.Data["json"] = string(response_put)
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *CulturaController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	if uid == "" {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	} else {
		err := models.DeleteCultura(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al eliminar cultura"}
		} else {
			u.Data["json"] = "delete success!"
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username	query 	string	true		"The username for login"
// @Param	password	query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *CulturaController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")

	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *CulturaController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
