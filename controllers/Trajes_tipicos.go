package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/models"
	"github.com/sena_2824182/pija_music_mid/services"
)

// Controller para usuarios
type Trajes_tipicosController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *Trajes_tipicosController) Post() {
	var body_Trajes map[string]interface{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Trajes); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Procesar el usuario
	json_Trajes_byte, err := json.Marshal(body_Trajes)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_Trajes, err := services.Metodo_post("services_post", json_Trajes_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error en el servicio POST"}
		u.ServeJSON()
		return
	}
	fmt.Println("Respuesta del POST:", string(response_Trajes))

	var Trajes_response map[string]interface{}
	if err := json.Unmarshal(response_Trajes, &Trajes_response); err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar la respuesta"}
		u.ServeJSON()
		return
	}

	u.Data["json"] = Trajes_response
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *Trajes_tipicosController) GetAll() {
	users := models.GetAllTrajes() // Corregido: ahora solo recibe un valor

	if len(users) == 0 { // Verifica si la lista está vacía
		u.Data["json"] = map[string]string{"error": "No hay trajes registrados"}
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
func (u *Trajes_tipicosController) Get() {
	uid := u.Ctx.Input.Param(":uid")
	if uid != "" {
		user, err := models.GetTrajes(uid)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al obtener trajes"}
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
func (u *Trajes_tipicosController) Put() {
	id_Trajes := u.Ctx.Input.Param(":uid")
	var body_Trajes map[string]interface{}

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &body_Trajes); err != nil {
		u.Data["json"] = map[string]string{"error": "Invalid input"}
		u.ServeJSON()
		return
	}

	// Obtener el usuario actual
	get_Trajes, err := services.Metodo_get("servicio", id_Trajes)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al obtener trajes"}
		u.ServeJSON()
		return
	}
	fmt.Println("Usuario obtenido:", string(get_Trajes))

	json_Trajes, err := services.ProcesadorJson(get_Trajes)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al procesar trajes"}
		u.ServeJSON()
		return
	}
	fmt.Println("JSON del usuario:", json_Trajes)

	// Construir nuevo JSON
	Trajes_tipicosActualizado := models.Trajes_tipicos{
		ID:       id_Trajes,
		Nombre_Trajes:   fmt.Sprintf("%v", body_Trajes["nombre_trajes"]),
		Region:   fmt.Sprintf("%v", body_Trajes["Region"]),
		Ocasion_Uso: fmt.Sprintf("%v", body_Trajes["Ocasion_Uso"]), 
		Materiales: fmt.Sprintf("%v",body_Trajes["Materiles"]),
		Descripcion_Trajes: fmt.Sprintf("%v",body_Trajes["Descripcion_Trajes"]),
		Imagen_Video:fmt.Sprintf("%v",body_Trajes["imagen_video"]),// Asegúrate de validar esto


	}
	
	err = models.UpdateTrajes(id_Trajes, Trajes_tipicosActualizado)
if err != nil {
    u.Data["json"] = map[string]string{"error": "Error al actualizar trajes"}
    u.ServeJSON()
    return
}

	// Actualizar usuario
	json_byte, err := json.Marshal(Trajes_tipicosActualizado)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al convertir a JSON"}
		u.ServeJSON()
		return
	}

	response_put, err := services.Metodo_put("servicio", id_Trajes, json_byte)
	if err != nil {
		u.Data["json"] = map[string]string{"error": "Error al actualizar trajes"}
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
func (u *Trajes_tipicosController) Delete() {
	id := u.Ctx.Input.Param(":uid")
	if id == "" {
		u.Data["json"] = map[string]string{"error": "User ID is required"}
	} else {
		err := models.DeleteTrajes(id)
		if err != nil {
			u.Data["json"] = map[string]string{"error": "Error al eliminar trajes"}
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
func (u *Trajes_tipicosController) Login() {
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
func (u *Trajes_tipicosController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
