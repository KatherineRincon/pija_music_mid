package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"

	
)

// FavoritosController operations for Favoritos
type FavoritosController struct {
	beego.Controller
}

// URLMapping ...
func (c *FavoritosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Favoritos
// @Param	body		body 	models.Favoritos	true		"body for Favoritos content"
// @Success 201 {object} models.Favoritos
// @Failure 403 body is empty
// @router / [post]
func (c *FavoritosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Favoritos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Favoritos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FavoritosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Favoritos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Favoritos
// @Failure 403
// @router / [get]
func (c *FavoritosController) GetAll() {
	fmt.Println("get all de favoritos")

		var Json_total_areglado []map[string]interface{}

	// ✅ Obtener el ID del usuario desde query param
	idUsuario := c.GetString("idUsuario")
    fmt.Println("idUsuario recibido:", idUsuario)

	if idUsuario == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Falta el parámetro idUsuario",
		}
		c.ServeJSON()
		return
	}

	// ✅ Obtener favoritos del usuario
	favoritos_bytes, err := services.Metodo_get("host_api", "Favoritos?limit=0")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los favoritos",
		}
		c.ServeJSON()
		return
	}

	jsonFavoritos, err := services.ProcesarJson(favoritos_bytes)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar los favoritos",
		}
		c.ServeJSON()
		return
	}

	favoritosSlice, _ := services.ConvertToSliceMap(jsonFavoritos["Data"])
	mapaFavoritos := make(map[interface{}]bool)

	for _, fav := range favoritosSlice {
		usuario, ok := fav["IdUsuario"].(map[string]interface{})
		if !ok {
			continue
		}
		if fmt.Sprintf("%v", usuario["Id"]) == idUsuario {
			idCancion := fav["IdCanciones"].(map[string]interface{})["Id"]
			mapaFavoritos[idCancion] = true
		}
	}

	// ✅ Obtener todas las canciones
	canciones_byte, err := services.Metodo_get("host_api", "Canciones?limit=0")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener las canciones",
		}
		c.ServeJSON()
		return
	}

	JsonCanciones, err := services.ProcesarJson(canciones_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar las canciones",
		}
		c.ServeJSON()
		return
	}

	Canciones_arreglo, _ := services.ConvertToSliceMap(JsonCanciones["Data"])
	grouped := make(map[string][]map[string]interface{})

	for _, cancion := range Canciones_arreglo {
		id := cancion["Id"]
		if !mapaFavoritos[id] {
			continue
		}

		if idArtistas, ok := cancion["IdArtistas"].(map[string]interface{}); ok {
			if nombreArtistico, ok := idArtistas["NombreArtistico"].(string); ok {
				grouped[nombreArtistico] = append(grouped[nombreArtistico], cancion)
			}
		}
	}

	// ✅ Armar respuesta final
	for nombreArtistico, canciones := range grouped {
		var arreglo_canciones_temporal []map[string]interface{}

		for _, cancion := range canciones {
			json_cancion := map[string]interface{}{
				"IdUsuario":        idUsuario,
				"IdCanciones":      cancion["Id"],
				"nombre":           cancion["TituloCancion"],
				"duracion":         cancion["Duracion"],
				"videoUrl":         cancion["RutaArchivo"],
				"Album":            cancion["Album"],
				"EstiloMusical":    cancion["EstiloMusical"],
				"FechaLanzamiento": cancion["FechaLanzamiento"],
			}
			arreglo_canciones_temporal = append(arreglo_canciones_temporal, json_cancion)
		}

		json_artista := map[string]interface{}{
			"nombre":    nombreArtistico,
			"imagen":    canciones[0]["IdArtistas"].(map[string]interface{})["ImagenVideo"],
			"canciones": arreglo_canciones_temporal,
		}

		Json_total_areglado = append(Json_total_areglado, json_artista)
	}

	// ✅ Respuesta final
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta de Canciones Favoritas del Usuario",
		"Data":    Json_total_areglado,
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Favoritos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Favoritos	true		"body for Favoritos content"
// @Success 200 {object} models.Favoritos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FavoritosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Favoritos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FavoritosController) Delete() {

}
