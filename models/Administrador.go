package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

// Administrador representa la estructura de un administrador en la base de datos
type Administrador struct {
    ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
    Nombre   string        `bson:"nombre" json:"nombre"`
    Correo   string        `bson:"correo" json:"correo"`
    Password string        `bson:"password" json:"password"`
}


var administradores = make(map[string]Administrador)

// AddOne creates a new Administrador and stores it in the map
func CreateAdmin(admin Administrador) string {
    admin.ID = bson.NewObjectId() // Asigna un nuevo ObjectId
    administradores[admin.ID.Hex()] = admin // Guarda el administrador con ID convertido a string
    return admin.ID.Hex() // Devuelve el ID como string
}


// GetOne retrieves an Administrador by ID
func GetAdministrador(id string) (Administrador, error) {
    if admin, ok := administradores[id]; ok {
        return admin, nil
    }
    return Administrador{}, errors.New("Administrador no encontrado")
}


// GetAll retrieves all Administradores
func GetAllAdmins() []Administrador {
    var list []Administrador
    for _, admin := range administradores {
        list = append(list, admin)
    }
    return list
}


// Update modifies an existing Administrador
func UpdateAdministrador(id string, admin Administrador) error {
	if _, ok := administradores[id]; !ok {
		return errors.New("Administrador not found")
	}
	admin.ID = bson.ObjectIdHex(id)
	administradores[id] = admin
	return nil
}

// Delete removes an Administrador
// Delete removes an Administrador
func DeleteAdministrador(id string) error {
    if _, ok := administradores[id]; !ok {
        return errors.New("Administrador not found")
    }
    
    delete(administradores, id) // Elimina el administrador del mapa

    return nil
}