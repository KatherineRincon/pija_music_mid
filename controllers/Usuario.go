package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/sena_2824182/pija_music_mid/services"
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
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_usuario, err := services.Metodo_post("services_post", json_usuario_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error en el servicio POST"}
		u.ServeJSON()
		return
	}
	fmt.Println("Respuesta del POST:", string(response_usuario))

	var usuario_response map[string]interface{}
	if err := json.Unmarshal(response_usuario, &usuario_response); err != nil {
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
func (u *UsuarioController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.GetUsuario(uid)
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
	get_usuario, err := services.Metodo_get("servicio", id_usuario)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al obtener usuario"}
		u.ServeJSON()
		return
	}
	fmt.Println("Usuario obtenido:", string(get_usuario))

	json_usuario, err := services.ProcesadorJson(get_usuario)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar usuario"}
		u.ServeJSON()
		return
	}
	fmt.Println("JSON del usuario:", json_usuario)

	// Construir nuevo JSON
	usuarioActualizado := models.Usuario{
		ID:       id_usuario,
		Nombre:   fmt.Sprintf("%v", body_usuario["nombre"]),
		Correo:   fmt.Sprintf("%v", body_usuario["email"]),
		Password: fmt.Sprintf("%v", body_usuario["password"]), // Asegúrate de validar esto
	}

	err = models.UpdateUsuario(id_usuario, usuarioActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar usuario"}
		u.ServeJSON()
		return
	}

	// Actualizar usuario
	json_byte, err := json.Marshal(usuarioActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_put, err := services.Metodo_put("servicio", id_usuario, json_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar usuario"}
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
func (u *UsuarioController) Delete() {
	uid := u.Ctx.Input.Param(":uid")
	if uid == "" {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	} else {
		err := models.DeleteUsuario(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al eliminar usuario"}
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
