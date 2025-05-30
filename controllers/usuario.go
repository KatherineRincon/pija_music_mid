package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"
)

// UsuarioController operations for Usuario
type UsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Postlogin", c.Postlogin)
}

// Postlogin ...
// @Title Create
// @Description create Usuario
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {object} models.Usuario
// @Failure 403 body is empty
// @router /login [post]
func (c *UsuarioController) Postlogin() {
	fmt.Println("postlogin")

	var body_ingresa map[string]interface{}

	var resultado map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("json ingresa", body_ingresa)
	}
	
	Email:= body_ingresa["Email"]
	Contraseña_cliente:= body_ingresa ["Contraseña"]

	Contraseña_cliente_string := fmt.Sprintf("%v", Contraseña_cliente)

	Email_string := fmt.Sprintf("%v", Email)

	endpoint:="Usuario?query=Email:" + Email_string
	
	
	body_usuario_byte,_:= services.Metodo_get("hots_crud", endpoint)


	body_usuario_json,_:= services.ProcesarJson(body_usuario_byte)

	data_usuario:= body_usuario_json["Data"]
	fmt.Println("data", data_usuario)
	if data_usuario==nil{
		fmt.Println("es nulo")

		resultado= map[string]interface{}{
			"Message": "usuario no existente",
			"Caso": 1,
		}

	} else {
		arreglo_usuario,_:= services.ConvertToSliceOfMaps(data_usuario)
		fmt.Println("Arreglo usuario", arreglo_usuario)
		contraseña_usuario:= arreglo_usuario[0]["IdCredenciales"].(map[string]interface{})["Contrasena"]
		fmt.Println("contraseña usuario", contraseña_usuario)
		Contraseña_string := fmt.Sprintf("%v", contraseña_usuario)

		if Contraseña_string== Contraseña_cliente_string {
			resultado= map[string]interface{}{
				"message": "usuario y correo correctos",
				"Caso": 2,
			}
		} else {
			resultado = map[string]interface{}{
				"message": "correo existente y contraseña no coincide",
				"Caso": 3,
			}
		}
		
}

		c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Consulta existosa",
		"Data":    resultado,
	}

	c.ServeJSON()

}

// Post ...
// @Title Create
// @Description create Usuario
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {object} models.Usuario
// @Failure 403 body is empty
// @router / [post]
func (c *UsuarioController) Post() {
	fmt.Println("post")

	var body_ingresa map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		fmt.Println("json ingresa", body_ingresa)
		fmt.Println("error", err)
	}

	body_contrasena := map[string]interface{}{
		"Contrasena": body_ingresa["Contraseña"],
	}

	bytes_contrasena, err := json.Marshal(body_contrasena)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_contrasena, _ := services.Metodo_post("hots_crud", "Credenciales", bytes_contrasena)

	var response_json_contrasena map[string]interface{}

	err1 := json.Unmarshal(body_response_contrasena, &response_json_contrasena)
	if err1 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	Id_contrasena := response_json_contrasena["Data"]
	Id_contrasena = Id_contrasena.(map[string]interface{})["Id"]

	fmt.Println("este es el id", Id_contrasena)

	fmt.Println("tipo dato", reflect.TypeOf(Id_contrasena))

	Id_entero := fmt.Sprintf("%v", Id_contrasena)

	fmt.Println("variable en string", Id_entero)

	Id, _ := strconv.Atoi(Id_entero)

	body_Usuario := map[string]interface{}{
		"Nombres":        body_ingresa["Nombres"],
		"Apellido":       body_ingresa["Apellido"],
		"Email":          body_ingresa["Email"],
		"IdCredenciales": map[string]interface{}{"Id": Id},
	}

	bytes_usuario, err := json.Marshal(body_Usuario)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	body_response_usuario, _ := services.Metodo_post("hots_crud", "Usuario", bytes_usuario)

	var response_json_usuario map[string]interface{}

	err2 := json.Unmarshal(body_response_usuario, &response_json_usuario)
	if err2 != nil {
		fmt.Println("Error al deserializar:", err)
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"Message": "Creación existosa",
		"Data":    body_ingresa,
	}
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Usuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuario
// @Failure 403
// @router / [get]
func (c *UsuarioController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Usuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuarioController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Usuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuarioController) Delete() {

}
