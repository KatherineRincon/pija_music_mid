package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Usuario representa la estructura de un usuario en la base de datos
type Trajes_tipicos struct {
	ID                 string `bson:"_id_Trajes,omitempty" json:"id_Trajes"`
	Nombre_Trajes      string `bson:"nombre_trajes" json:"nombre_trajes"`
	Region             string `bson:"region" json:"reguin"`
	Ocasion_Uso        string `bson:"ocasion" json:"ocasion"`
	Materiales         string `bson:"materiales"json:"materiles"`
	Descripcion_Trajes string `bson"descripcion_trajes"json:"descripcion_trajes"`
	Imagen_Video       string `bson"imagen_video"json:"imagen_video"`
}

// Mapa para almacenar trajes en memoria (esto simula una base de datos)
var Trajes = make(map[string]Trajes_tipicos)

func AddTrajes(user Trajes_tipicos) string {
	user.ID = bson.NewObjectId().Hex() // Generar un nuevo ID único
	Trajes[user.ID] = user          // Guardar el usuario en el mapa
	return user.ID                     // Retornar el ID generado
}

func GetTrajes(id string) (Trajes_tipicos, error) {
	if user, ok := Trajes[id]; ok {
		return user, nil
	}
	return Trajes_tipicos{}, errors.New("trajes no encontrado")
}

func GetAllTrajes() []Trajes_tipicos {
	var lista []Trajes_tipicos
	for _, user := range Trajes {
		lista = append(lista, user)
	}
	return lista
}

func UpdateTrajes(id string, user Trajes_tipicos) error {
	if _, ok := Trajes[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	user.ID = id // Mantener el mismo ID
	Trajes[id] = user
	return nil
}

func DeleteTrajes(id string) error {
	if _, ok := Trajes[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	delete(Trajes, id)
	return nil
}
