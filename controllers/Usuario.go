package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/astaxie/beego"
)

// Controller para usuarios
type UsuarioController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UsuarioController) Post() {
	var body_usuario map[string]interface{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_usuario); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Procesar el usuario
	json_usuario_byte, err := json.Marshal(body_usuario)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_usuario := services.Metodo_post("services_post", json_usuario_byte)
	fmt.Println("Respuesta del POST:", string(response_usuario))

	var usuario_response map[string]interface{}
	if err := json.Unmarshal(response_usuario, &usuario_response); err != nil {
		fmt.Println("Error al procesar la respuesta:", err)
		u.Data["json"] = map[string]string{"error": "Error al procesar la respuesta"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = usuario_response
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UsuarioController) GetAll() {
	users := models.GetAllUsuario()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UsuarioController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.GetUsuario(uid)
		if err != nil {
			u.Data["json"] = err.Error()
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
func (u *UsuarioController) Put() {
	id_usuario := u.Ctx.Input.Param(":uid")
	var body_usuario map[string]interface{}

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_usuario); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Obtener el usuario actual
	get_usuario, _ := services.Metodo_get("servicio", id_usuario)
	fmt.Println("Usuario obtenido:", string(get_usuario))

	json_usuario, _ := services.ProcesadorJson(get_usuario)
	fmt.Println("JSON del usuario:", json_usuario)

	// Construir nuevo JSON
	nuevo_json := map[string]interface{}{
		"id":       json_usuario["id"],
		"nombre":   body_usuario["nombre"],
		"email":    body_usuario["email"],
		"apellido": body_usuario["apellido"],
	}
	fmt.Println("Nuevo JSON:", nuevo_json)

	// Actualizar usuario
	json_byte, _ := json.Marshal(nuevo_json)
	response_put, _ := services.Metodo_put("servicio", id_usuario, json_byte)
	fmt.Println("Respuesta del PUT:", string(response_put))

	u.Data["json"] = response_put
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UsuarioController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	models.DeleteUsuario(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	Usuarioname		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UsuarioController) Login() {
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
func (u *UsuarioController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
