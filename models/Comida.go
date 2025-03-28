package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Usuario representa la estructura de un usuario en la base de datos
type Comida struct {
	ID_Comida          string `bson:"_id_Comida,omitempty" json:"id_Comido"`
	Nombre_comida      string `bson:"nombre_comida" json:"nombre_comida"`
	Tipo_comida        string `bson:"tipo_comida" json:"tipo_comida"`
	Description_comida string `bson:"description_comida" json:"description_comida"`
	Imagen_Video       string `bson:"imagen_video"json:"imagen_video"`
}

// Mapa para almacenar usuarios en memoria (esto simula una base de datos)
var comida = make(map[string]Comida)

func AddComida(user Comida) string {
	user.ID_Comida = bson.NewObjectId().Hex() // Generar un nuevo ID único
	comida[user.ID_Comida] = user             // Guardar el comida en el mapa
	return user.ID_Comida                     // Retornar el ID generado
}

func Getcomida(id string) (Comida, error) {
	if user, ok := comida[id]; ok {
		return user, nil
	}
	return Comida{}, errors.New("Comida no encontrado")
}

func GetAllComida() []Comida {
	var lista []Comida
	for _, user := range comida {
		lista = append(lista, user)
	}
	return lista
}

func UpdateComida(id string, user Comida) error {
	if _, ok := comida[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	user.ID_Comida = id // Mantener el mismo ID
	comida[id] = user
	return nil
}

func DeleteComida(id string) error {
	if _, ok := comida[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	delete(comida, id)
	return nil
}
