package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Administrador struct {
	ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Nombre   string        `bson:"nombre" json:"nombre"`
	Correo   string        `bson:"correo" json:"correo"`
	Password string        `bson:"password" json:"password"`
}

var administradores = make(map[string]Administrador)

// AddOne creates a new Administrador and stores it in the map
func AddOne(admin Administrador) string {
	admin.ID = bson.NewObjectId()
	administradores[admin.ID.Hex()] = admin
	return admin.ID.Hex()
}

// GetOne retrieves an Administrador by ID
func GetOne(id string) (Administrador, error) {
	if admin, ok := administradores[id]; ok {
		return admin, nil
	}
	return Administrador{}, errors.New("Administrador not found")
}

// GetAll retrieves all Administradores
func GetAll() []Administrador {
	var list []Administrador
	for _, admin := range administradores {
		list = append(list, admin)
	}
	return list
}

// Update modifies an existing Administrador
func Update(id string, admin Administrador) error {
	if _, ok := administradores[id]; !ok {
		return errors.New("Administrador not found")
	}
	admin.ID = bson.ObjectIdHex(id)
	administradores[id] = admin
	return nil
}

// Delete removes an Administrador
func Delete(id string) error {
	if _, ok := administradores[id]; !ok {
		return errors.New("Administrador not found")
	}
	delete(administradores, id)
	return nil
}
