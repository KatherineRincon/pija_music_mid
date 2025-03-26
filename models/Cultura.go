package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// Usuario representa la estructura de un usuario en la base de datos
type Cultura struct {
	ID      string `bson:"_id,omitempty" json:"id"`
	Nombre_Cultura   string `bson:"nombre_cultuta" json:"nombre_cultura"`
	Descripcion_Cultura   string `bson:"descpcion_cultura" json:"descpcion_cultura"`
	Imagen_video string `bson:"imagen_video" json:"imagen_video"`
}

// Mapa para almacenar usuarios en memoria (esto simula una base de datos)
var cultura = make(map[string]Cultura)


func AddCultura(user Cultura) string {
	user.ID = bson.NewObjectId().Hex() // Generar un nuevo ID único
	cultura[user.ID] = user           // Guardar el usuario en el mapa
	return user.ID                      // Retornar el ID generado
}

func GetCultura(id string) (Cultura, error) {
	if user, ok := cultura[id]; ok {
		return user, nil
	}
	return Cultura{}, errors.New("Usuario no encontrado")
}

func GetAllCultura() []Cultura {
	var lista []Cultura
	for _, user := range cultura {
		lista = append(lista, user)
	}
	return lista
}

func UpdateCultura(id string, user Cultura) error {
	if _, ok := cultura[id]; !ok {
		return errors.New("Cultura no encontrado")
	}
	user.ID = id // Mantener el mismo ID
	cultura[id] = user
	return nil
}

func DeleteCultura(id string) error {
	if _, ok := cultura[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	delete(cultura, id)
	return nil
}
