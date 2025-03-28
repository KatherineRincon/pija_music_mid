package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/sena_2824182/pija_music_mid/services"
)

// Controller para usuarios
type ComidaController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *ComidaController) Post() {
	var body_comida map[string]interface{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_comida); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Procesar el usuario
	json_comida_byte, err := json.Marshal(body_comida)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_comida, err := services.Metodo_post("services_post", json_comida_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error en el servicio POST"}
		u.ServeJSON()
		return
	}
	fmt.Println("Respuesta del POST:", string(response_comida))

	var comida_response map[string]interface{}
	if err := json.Unmarshal(response_comida, &comida_response); err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar la respuesta"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = comida_response
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *ComidaController) GetAll() {
	users := models.GetAllUsuario() // Corregido: ahora solo recibe un valor

	if len(users) == 0 { // Verifica si la lista está vacía
		u.Data["json"] = map[string]string{"error": "No hay usuarios registrados"}
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
func (u *ComidaController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.Getcomida(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al obtener comida"}
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
func (u *ComidaController) Put() {
	id_Comida := u.Ctx.Input.Param(":uid")
	var body_comida map[string]interface{}

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_comida); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Obtener el comida actual
	get_comida, err := services.Metodo_get("servicio", 	id_Comida)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al obtener comida"}
		u.ServeJSON()
		return
	}
	fmt.Println("comida obtenido:", string(get_comida))

	json_comida, err := services.ProcesadorJson(get_comida)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar comida"}
		u.ServeJSON()
		return
	}
	fmt.Println("JSON del comida:", json_comida)

	// Construir nuevo JSON
	ComidaActualizado := models.Comida{
		ID_Comida:          id_Comida,
		Nombre_comida:      fmt.Sprintf("%v", body_comida["nombre_comida"]),
		Tipo_comida:        fmt.Sprintf("%v", body_comida["tipo_comida"]),
		Description_comida: fmt.Sprintf("%v", body_comida["description_comida"]), // Asegúrate de validar esto
		Imagen_Video:       fmt.Sprintf("%v", body_comida["imagen_video"]),
	}

	err = models.UpdateComida(id_Comida, ComidaActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar comida"}
		u.ServeJSON()
		return
	}

	// Actualizar comida
	json_byte, err := json.Marshal(ComidaActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_put, err := services.Metodo_put("servicio", id_Comida, json_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar comida"}
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
func (u *ComidaController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	if uid == "" {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	} else {
		err := models.DeleteComida(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al eliminar comida"}
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
func (u *ComidaController) Login() {
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
func (u *ComidaController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
