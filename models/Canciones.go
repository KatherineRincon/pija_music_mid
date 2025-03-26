package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	
	"gopkg.in/mgo.v2"
)

// Canciones representa la estructura de una canción en la base de datos
type Canciones struct {
	IdCanciones     string        `bson:"_id,omitempty" json:"idcanciones"`
	TituloCancion   string        `bson:"titulo" json:"titulocancion"`
	IdArtista       string        `bson:"idartista" json:"idartista"`
	Genero          string        `bson:"genero" json:"genero"`
	Album           string        `bson:"album" json:"album"`
	IdEstilo        string        `bson:"idestilo" json:"idestilo"`
	FechaLanzamiento string       `bson:"fechalanzamiento" json:"fechalanzamiento"`
	Duracion        string        `bson:"duracion" json:"duracion"`
	RutaArchivo     string        `bson:"rutaarchivo" json:"rutaarchivo"`
	Activo          bool          `bson:"activo" json:"activo"`
	FechaCreacion   string        `bson:"fechacreacion" json:"fechacreacion"`
	FechaModificacion string      `bson:"fechamodificacion" json:"fechamodificacion"`
}
var (
	dbName       = "pija_music"
	collection   = "canciones"
	mongoSession *mgo.Session
)

// init inicializa la conexión con MongoDB
func init() {
	var err error
	mongoSession, err = mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
}

// CreateCanciones guarda una nueva canción en la base de datos
func CreateCanciones(cancion Canciones) (*Canciones, error) {
	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB(dbName).C(collection)
	cancion.ID = bson.NewObjectId()

	if err := c.Insert(cancion); err != nil {
		return nil, err
	}
	return &cancion, nil
}

// GetCanciones obtiene una canción por su ID
func GetCanciones(id string) (*Canciones, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("Invalid ID format")
	}

	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB(dbName).C(collection)
	var cancion Canciones
	if err := c.FindId(bson.ObjectIdHex(id)).One(&cancion); err != nil {
		return nil, err
	}
	return &cancion, nil
}

// GetAllCanciones obtiene todas las canciones
func GetAllCanciones() ([]Canciones, error) {
	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB(dbName).C(collection)
	var canciones []Canciones
	if err := c.Find(nil).All(&canciones); err != nil {
		return nil, err
	}
	return canciones, nil
}

// UpdateCanciones actualiza una canción por su ID
func UpdateCanciones(id string, cancion Canciones) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid ID format")
	}

	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB(dbName).C(collection)
	return c.UpdateId(bson.ObjectIdHex(id), cancion)
}

// DeleteCanciones elimina una canción por su ID
func DeleteCanciones(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid ID format")
	}

	session := mongoSession.Copy()
	defer session.Close()

	c := session.DB(dbName).C(collection)
	return c.RemoveId(bson.ObjectIdHex(id))
}