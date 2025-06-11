package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/pija_music_mid/services"
)

// CancionesController operations for Canciones
type CancionesController struct {
	beego.Controller
}

type IdArtista struct {
	Activo            bool   `json:"Activo"`
	Biografia         string `json:"Biografia"`
	FechaCreacion     string `json:"FechaCreacion"`
	FechaModificacion string `json:"FechaModificacion"`
	Id                int    `json:"Id"`
	ImagenVideo       string `json:"ImagenVideo"`
	NombreArtistico   string `json:"NombreArtistico"`
	NombreReal        string `json:"NombreReal"`
	RedesSociales     string `json:"RedesSociales"`
}

type IdEstilo struct {
	Activo                  bool   `json:"Activo"`
	DescripcionMusical      string `json:"DescripcionMusical"`
	FechaCreacion           string `json:"FechaCreacion"`
	FechaModificacion       string `json:"FechaModificacion"`
	Id                      int    `json:"Id"`
	InstrumentosPrincipales string `json:"InstrumentosPrincipales"`
	NombreGenero            string `json:"NombreGenero"`
}

type Cancion struct {
	Activo            bool      `json:"Activo"`
	Album             string    `json:"Album"`
	Duracion          string    `json:"Duracion"`
	FechaCreacion     string    `json:"FechaCreacion"`
	FechaLanzamiento  string    `json:"FechaLanzamiento"`
	FechaModificacion string    `json:"FechaModificacion"`
	Id                int       `json:"Id"`
	IdArtistas        IdArtista `json:"IdArtistas"`
	IdEstilo          IdEstilo  `json:"IdEstilo"`
	RutaArchivo       string    `json:"RutaArchivo"`
	TituloCancion     string    `json:"TituloCancion"`
}

// URLMapping ...
func (c *CancionesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Canciones
// @Param	body		body 	models.Canciones	true		"body for Canciones content"
// @Success 201 {object} models.Canciones
// @Failure 403 body is empty
// @router / [post]
func (c *CancionesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Canciones by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Canciones
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CancionesController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Canciones
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Canciones
// @Failure 403
// @router / [get]
func (c *CancionesController) GetAll() {
	fmt.Println("get all de canciones")

	var json_arreglado map[string]interface{}
	var Json_total_areglado []map[string]interface{}
	var arreglo_canciones_temporal []map[string]interface{}
	var json_cancion_temporal map[string]interface{}

	canciones_byte, _ := services.Metodo_get("host_api", "Canciones?limit=0")
	JsonCanciones, _ := services.ProcesarJson(canciones_byte)
	//fmt.Println("Json Canciones",JsonCanciones)

	SoloCanciones := JsonCanciones["Data"]
	//fmt.Println("Solo Canciones", SoloCanciones)

	Canciones_arreglo, _ := services.ConvertToSliceMap(SoloCanciones)

	///////////////////////////////////////////////////////////////////////

	grouped := make(map[string][]map[string]interface{})

	// Agrupar canciones
	for _, cancion := range Canciones_arreglo {
		// Obtener el NombreArtístico desde el campo IdArtistas.NombreArtistico
		if idArtistas, ok := cancion["IdArtistas"].(map[string]interface{}); ok {
			if nombreArtistico, ok := idArtistas["NombreArtistico"].(string); ok {
				// Guardar la canción en el grupo correspondiente
				grouped[nombreArtistico] = append(grouped[nombreArtistico], cancion)
			}
		}
	}

	fmt.Println("agrupacion de canciones", grouped)

	// Mostrar las canciones agrupadas
	for nombreArtistico, canciones := range grouped {
    arreglo_canciones_temporal = nil

    for _, cancion := range canciones {
        fmt.Printf("  Titulo: %s, Album: %s\n", cancion["TituloCancion"], cancion["Album"])

        json_cancion_temporal = map[string]interface{}{
            "IdCanciones": cancion["IdCanciones"],
            "nombre":    cancion["TituloCancion"],
            "duracion":  cancion["Duracion"],
            "videoUrl":  cancion["RutaArchivo"],
        }

        arreglo_canciones_temporal = append(arreglo_canciones_temporal, json_cancion_temporal)
    }

    // ✅ Esta parte debe ir fuera del bucle de canciones
    json_arreglado = map[string]interface{}{
        "nombre":    nombreArtistico,
        "imagen":    canciones[0]["IdArtistas"].(map[string]interface{})["ImagenVideo"],
        "canciones": arreglo_canciones_temporal,
    }

    Json_total_areglado = append(Json_total_areglado, json_arreglado)
}


	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200, 
		"Message": "Consulta de Canciones",
		"Data":    Json_total_areglado,
	}

	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Canciones
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Canciones	true		"body for Canciones content"
// @Success 200 {object} models.Canciones
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CancionesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Canciones
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CancionesController) Delete() {

}
