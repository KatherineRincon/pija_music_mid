package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"
)

// LugaresController operations for Lugares
type LugaresController struct {
	beego.Controller
}

// URLMapping ...
func (c *LugaresController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

}

// Post ...
// @Title Create
// @Description create Lugares
// @Param	body		body 	models.Lugares	true		"body for Lugares content"
// @Success 201 {object} models.Lugares
// @Failure 403 body is empty
// @router / [post]
func (c *LugaresController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Lugares by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Lugares
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LugaresController) GetOne() {
	fmt.Println("one")

}

// GetAll ...
// @Title GetAll
// @Description get Lugares
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Lugares
// @Failure 403
// @router / [get]
func (c *LugaresController) GetAll() {

	/// Vatriable generales
	var lugares_organizados_totales map[string]interface{}
	var lugares_organizados_senderismo []map[string]interface{}
	var lugares_senderismo_temporal map[string]interface{}
	var lugares_organizados_actividad []map[string]interface{}
	var lugares_actividad_temporal map[string]interface{}
	var lugares_organizados_historicos []map[string]interface{}
	var lugares_historicos_temporal map[string]interface{}
	var lugares_organizados_culturales []map[string]interface{}
	var lugares_culturales_temporal map[string]interface{}
	var lugares_organizados_restaurante []map[string]interface{}
	var lugares_restaurante_temporal map[string]interface{}

	/// este el flujo para senderismo

	body_lugares_senderismo_byte, _ := services.Metodo_get("host_api", "Lugares?limit=0&query=IdTipoLugares.Id:1")
	body_lugares__senderismo_json, _ := services.ProcesarJson(body_lugares_senderismo_byte)
	//fmt.Println(body_lugares__senderismo_json)

	lugares_senderismo := body_lugares__senderismo_json["Data"]

	lugares_senderismo_arreglo, _ := services.ConvertToSliceMap(lugares_senderismo)

	for _, lugar := range lugares_senderismo_arreglo {

		//fmt.Println("Lugar numero", i)
		//fmt.Println(lugar)

		// Parte que saca el json de las imagenes en partes, ejemplo imagen, imagenes y fondo

		imagenVideoStr, _ := lugar["ImagenVideo"].(string)
		var imagenVideoData map[string]interface{}
		err := json.Unmarshal([]byte(imagenVideoStr), &imagenVideoData)
		if err != nil {
			fmt.Println("Error al parsear ImagenVideo:", err)
			return
		}

		imagenesClean := strings.ReplaceAll(imagenVideoData["imagenes"].(string), "'", "\"")
		var imagenes []string
		err = json.Unmarshal([]byte(imagenesClean), &imagenes)

		descripcionClean := strings.ReplaceAll(lugar["DescripcionLugares"].(string), "'", "\"")

		descripcionJSON := "[" + descripcionClean + "]"
		descripcionJSON = strings.ReplaceAll(descripcionJSON, "'", "\"")

		re := regexp.MustCompile(`([{,]\s*)(\w+)(\s*:)`)
		descripcionJSON = re.ReplaceAllString(descripcionJSON, `$1"$2"$3`)
		var descripcion []map[string]string
		err = json.Unmarshal([]byte(descripcionJSON), &descripcion)
		if err != nil {
			fmt.Println("Contenido problemático:", descripcionClean)
			fmt.Println("Error al parsear la descripción:", err)
			return
		}

		lugares_senderismo_temporal = map[string]interface{}{
			"titulo": lugar["NombreLugar"],
			"imagen": imagenVideoData["imagen"],
			"texto":  lugar["Direccion"],

			"detalle": map[string]interface{}{
				"fondo":       imagenVideoData["fondo"],
				"imagenes":    imagenes,
				"titulo":      lugar["NombreLugar"],
				"descripcion": descripcion,
			},
		}

		lugares_organizados_senderismo = append(lugares_organizados_senderismo, lugares_senderismo_temporal)

	}

	//Inicia el flujo de actividad

	body_lugares_actividad_byte, _ := services.Metodo_get("host_api", "Lugares?limit=0&query=IdTipoLugares.Id:2")
	body_lugares__actividad_json, _ := services.ProcesarJson(body_lugares_actividad_byte)
	//fmt.Println(body_lugares__senderismo_json)

	lugares_actividad := body_lugares__actividad_json["Data"]

	lugares_actividad_arreglo, _ := services.ConvertToSliceMap(lugares_actividad)
	for _, lugar := range lugares_actividad_arreglo {

		// Parte que saca el json de las imagenes en partes, ejemplo imagen, imagenes y fondo

		imagenVideoStr, _ := lugar["ImagenVideo"].(string)
		var imagenVideoData map[string]interface{}
		err := json.Unmarshal([]byte(imagenVideoStr), &imagenVideoData)
		if err != nil {
			fmt.Println("Error al parsear ImagenVideo:", err)
			return
		}

		imagenesClean := strings.ReplaceAll(imagenVideoData["imagenes"].(string), "'", "\"")
		var imagenes []string
		err = json.Unmarshal([]byte(imagenesClean), &imagenes)

		descripcionClean := strings.ReplaceAll(lugar["DescripcionLugares"].(string), "'", "\"")

		descripcionJSON := "[" + descripcionClean + "]"
		descripcionJSON = strings.ReplaceAll(descripcionJSON, "'", "\"")

		re := regexp.MustCompile(`([{,]\s*)(\w+)(\s*:)`)
		descripcionJSON = re.ReplaceAllString(descripcionJSON, `$1"$2"$3`)
		//var descripcion []map[string]string

		descripcionJSON = strings.ReplaceAll(descripcionJSON, ",]", "]")

		descripcionStr := strings.ReplaceAll(descripcionJSON, "\n", " ")

		// 2. Parsear directamente
		var descripcion []map[string]string
		err = json.Unmarshal([]byte(descripcionStr), &descripcion)
		if err != nil {
			fmt.Println("Error al parsear la descripción:", err)
			return
		}

		lugares_actividad_temporal = map[string]interface{}{
			"titulo":   lugar["NombreLugar"],
			"imagen":   imagenVideoData["imagen"],
			"texto":    lugar["Direccion"],
			"id_lugar": lugar["Id"],

			"detalle": map[string]interface{}{
				"fondo":       imagenVideoData["fondo"],
				"imagenes":    imagenes,
				"titulo":      lugar["NombreLugar"],
				"descripcion": descripcion,
			},
		}

		lugares_organizados_actividad = append(lugares_organizados_actividad, lugares_actividad_temporal)

	}

	//Inicia el flujo de Lugares historicos

	body_lugares_historia_byte, _ := services.Metodo_get("host_api", "Lugares?limit=0&query=IdTipoLugares.Id:3")
	body_lugares__historia_json, _ := services.ProcesarJson(body_lugares_historia_byte)

	lugares_historia := body_lugares__historia_json["Data"]

	lugares_historia_arreglo, _ := services.ConvertToSliceMap(lugares_historia)

	for _, lugar := range lugares_historia_arreglo {

		//fmt.Println("Lugar numero", i)
		//fmt.Println(lugar)

		// Parte que saca el json de las imagenes en partes, ejemplo imagen, imagenes y fondo

		imagenVideoStr, _ := lugar["ImagenVideo"].(string)
		var imagenVideoData map[string]interface{}
		err := json.Unmarshal([]byte(imagenVideoStr), &imagenVideoData)
		if err != nil {
			fmt.Println("Error al parsear ImagenVideo:", err)
			return
		}

		imagenesClean := strings.ReplaceAll(imagenVideoData["imagenes"].(string), "'", "\"")
		var imagenes []string
		err = json.Unmarshal([]byte(imagenesClean), &imagenes)

		descripcionClean := strings.ReplaceAll(lugar["DescripcionLugares"].(string), "'", "\"")

		descripcionJSON := "[" + descripcionClean + "]"
		descripcionJSON = strings.ReplaceAll(descripcionJSON, "'", "\"")

		re := regexp.MustCompile(`([{,]\s*)(\w+)(\s*:)`)
		descripcionJSON = re.ReplaceAllString(descripcionJSON, `$1"$2"$3`)

		descripcionJSON = strings.ReplaceAll(descripcionJSON, ",]", "]")

		descripcionStr := strings.ReplaceAll(descripcionJSON, "\n", " ")

		// 2. Parsear directamente
		var descripcion []map[string]string
		err = json.Unmarshal([]byte(descripcionStr), &descripcion)
		if err != nil {
			fmt.Println("Error al parsear la descripción:", err)
			return
		}

		lugares_historicos_temporal = map[string]interface{}{
			"titulo": lugar["NombreLugar"],
			"imagen": imagenVideoData["imagen"],
			"texto":  lugar["Direccion"],

			"detalle": map[string]interface{}{
				"fondo":       imagenVideoData["fondo"],
				"imagenes":    imagenes,
				"titulo":      lugar["NombreLugar"],
				"descripcion": descripcion,
			},
		}

		lugares_organizados_historicos = append(lugares_organizados_historicos, lugares_historicos_temporal)

	}

	// Inicia flujo culturales

	body_lugares_culturales_byte, _ := services.Metodo_get("host_api", "Lugares?limit=0&query=IdTipoLugares.Id:4")
	body_lugares__culturales_json, _ := services.ProcesarJson(body_lugares_culturales_byte)
	//fmt.Println(body_lugares__senderismo_json)

	lugares_culturales := body_lugares__culturales_json["Data"]

	lugares_culturales_arreglo, _ := services.ConvertToSliceMap(lugares_culturales)

	for _, lugar := range lugares_culturales_arreglo {

		//fmt.Println("Lugar numero", i)
		//fmt.Println(lugar)

		// Parte que saca el json de las imagenes en partes, ejemplo imagen, imagenes y fondo

		imagenVideoStr, _ := lugar["ImagenVideo"].(string)
		var imagenVideoData map[string]interface{}
		err := json.Unmarshal([]byte(imagenVideoStr), &imagenVideoData)
		if err != nil {
			fmt.Println("Error al parsear ImagenVideo:", err)
			return
		}

		imagenesClean := strings.ReplaceAll(imagenVideoData["imagenes"].(string), "'", "\"")
		var imagenes []string
		err = json.Unmarshal([]byte(imagenesClean), &imagenes)

		descripcionClean := strings.ReplaceAll(lugar["DescripcionLugares"].(string), "'", "\"")

		descripcionJSON := "[" + descripcionClean + "]"
		descripcionJSON = strings.ReplaceAll(descripcionJSON, "'", "\"")

		re := regexp.MustCompile(`([{,]\s*)(\w+)(\s*:)`)

		descripcionJSON = re.ReplaceAllString(descripcionJSON, `$1"$2"$3`)
		descripcionJSON = strings.ReplaceAll(descripcionJSON, ",]", "]")
		descripcionStr := strings.ReplaceAll(descripcionJSON, "\n", " ")

		var descripcion []map[string]string
		err = json.Unmarshal([]byte(descripcionStr), &descripcion)
		if err != nil {
			fmt.Println("Contenido problemático:", descripcionClean)
			fmt.Println("Error al parsear la descripción:", err)
			return
		}

		lugares_culturales_temporal = map[string]interface{}{
			"titulo": lugar["NombreLugar"],
			"imagen": imagenVideoData["imagen"],
			"texto":  lugar["Direccion"],

			"detalle": map[string]interface{}{
				"fondo":       imagenVideoData["fondo"],
				"imagenes":    imagenes,
				"titulo":      lugar["NombreLugar"],
				"descripcion": descripcion,
			},
		}

		lugares_organizados_culturales = append(lugares_organizados_culturales, lugares_culturales_temporal)

	}

	// Inicia flujo restaurantes

	body_lugares_restaurante_byte, _ := services.Metodo_get("host_api", "Lugares?limit=0&query=IdTipoLugares.Id:5")
	body_lugares__restaurante_json, _ := services.ProcesarJson(body_lugares_restaurante_byte)
	//fmt.Println(body_lugares__senderismo_json)

	lugares_restaurante := body_lugares__restaurante_json["Data"]

	lugares_restaurante_arreglo, _ := services.ConvertToSliceMap(lugares_restaurante)

	for _, lugar := range lugares_restaurante_arreglo {

		//fmt.Println("Lugar numero", i)
		//fmt.Println(lugar)

		// Parte que saca el json de las imagenes en partes, ejemplo imagen, imagenes y fondo

		imagenVideoStr, _ := lugar["ImagenVideo"].(string)
		var imagenVideoData map[string]interface{}
		err := json.Unmarshal([]byte(imagenVideoStr), &imagenVideoData)
		if err != nil {
			fmt.Println("Error al parsear ImagenVideo:", err)
			return
		}

		imagenesClean := strings.ReplaceAll(imagenVideoData["imagenes"].(string), "'", "\"")
		var imagenes []string
		err = json.Unmarshal([]byte(imagenesClean), &imagenes)

		descripcionClean := strings.ReplaceAll(lugar["DescripcionLugares"].(string), "'", "\"")

		descripcionJSON := "[" + descripcionClean + "]"
		descripcionJSON = strings.ReplaceAll(descripcionJSON, "'", "\"")

		re := regexp.MustCompile(`([{,]\s*)(\w+)(\s*:)`)

		descripcionJSON = re.ReplaceAllString(descripcionJSON, `$1"$2"$3`)
		descripcionJSON = strings.ReplaceAll(descripcionJSON, ",]", "]")
		descripcionStr := strings.ReplaceAll(descripcionJSON, "\n", " ")

		var descripcion []map[string]string
		err = json.Unmarshal([]byte(descripcionStr), &descripcion)
		if err != nil {
			fmt.Println("Contenido problemático:", descripcionClean)
			fmt.Println("Error al parsear la descripción:", err)
			return
		}

		lugares_restaurante_temporal = map[string]interface{}{
			"titulo": lugar["NombreLugar"],
			"imagen": imagenVideoData["imagen"],
			"texto":  lugar["Direccion"],

			"detalle": map[string]interface{}{
				"fondo":       imagenVideoData["fondo"],
				"imagenes":    imagenes,
				"titulo":      lugar["NombreLugar"],
				"descripcion": descripcion,
			},
		}

		lugares_organizados_restaurante = append(lugares_organizados_restaurante, lugares_restaurante_temporal)

	}

	// Construcionm finsl

	lugares_organizados_totales = map[string]interface{}{
		"senderismo":         lugares_organizados_senderismo,
		"actividad_familiar": lugares_organizados_actividad,
		"lugares_historicos": lugares_organizados_historicos,
		"eventos_culturales": lugares_organizados_culturales,
		"restaurantes":       lugares_organizados_restaurante,
	}

	c.Data["json"] = map[string]interface{}{"Success": true,
		"Status":  200,
		"Message": "Actualización exitosa",
		"Data":    lugares_organizados_totales,
	}

	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Lugares
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Lugares	true		"body for Lugares content"
// @Success 200 {object} models.Lugares
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LugaresController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Lugares
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LugaresController) Delete() {

}
