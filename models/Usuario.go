package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// Usuario representa la estructura de un usuario en la base de datos
type Usuario struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Nombre   string `bson:"nombre" json:"nombre"`
	Correo   string `bson:"correo" json:"correo"`
	Password string `bson:"password" json:"password"`
}

// Mapa para almacenar usuarios en memoria (esto simula una base de datos)
var usuarios = make(map[string]Usuario)


func AddUsuario(user Usuario) string {
	user.ID = bson.NewObjectId().Hex() // Generar un nuevo ID único
	usuarios[user.ID] = user           // Guardar el usuario en el mapa
	return user.ID                      // Retornar el ID generado
}

func GetUsuario(id string) (Usuario, error) {
	if user, ok := usuarios[id]; ok {
		return user, nil
	}
	return Usuario{}, errors.New("Usuario no encontrado")
}

func GetAllUsuario() []Usuario {
	var lista []Usuario
	for _, user := range usuarios {
		lista = append(lista, user)
	}
	return lista
}

func UpdateUsuario(id string, user Usuario) error {
	if _, ok := usuarios[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	user.ID = id // Mantener el mismo ID
	usuarios[id] = user
	return nil
}

func DeleteUsuario(id string) error {
	if _, ok := usuarios[id]; !ok {
		return errors.New("Usuario no encontrado")
	}
	delete(usuarios, id)
	return nil
}

