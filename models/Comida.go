package models

import (
	"context"
	"errors"
	"time"

	
	

)

// Comida representa la estructura de una comida en la base de datos
type Comida struct {
	IdComida           string `bson:"_id,omitempty" json:"idcomida"`
	NombreComida       string `bson:"nombrecomida" json:"nombrecomida"`
	TipoComida         string `bson:"tipocomida" json:"tipocomida"`
	DescripcionComida  string `bson:"descripcioncomida" json:"descripcioncomida"`
	ImagenVideo        string `bson:"imagenvideo" json:"imagenvideo"`
	Activo             bool   `bson:"activo" json:"activo"`
	FechaCreacion      string `bson:"fechacreacion" json:"fechacreacion"`
	FechaModificacion  string `bson:"fechamodificacion" json:"fechamodificacion"`
}

var (
	client     *mongo.Client
	collection *mongo.Collection
	mongoSession *mgo.Session
)

// InitDB conecta a MongoDB y asigna la colección
func InitDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	collection = client.Database("pija_music").Collection("comidas")
	return nil
}

// CreateComida guarda una nueva comida en la base de datos
func CreateComida(comida Comida) (*Comida, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	comida.IdComida = uuid.New().String() // Genera un UUID como ID único
	comida.FechaCreacion = time.Now().Format(time.RFC3339)
	comida.FechaModificacion = time.Now().Format(time.RFC3339)

	_, err := collection.InsertOne(ctx, comida)
	if err != nil {
		return nil, err
	}
	return &comida, nil
}

// GetComida obtiene una comida por su ID
func GetComida(id string) (*Comida, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var comida Comida
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comida)
	if err != nil {
		return nil, errors.New("comida no encontrada")
	}
	return &comida, nil
}

// GetAllComidas obtiene todas las comidas
func GetAllComidas() ([]Comida, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var comidas []Comida
	for cursor.Next(ctx) {
		var comida Comida
		if err := cursor.Decode(&comida); err != nil {
			return nil, err
		}
		comidas = append(comidas, comida)
	}
	return comidas, nil
}

// UpdateComida actualiza una comida por su ID
func UpdateComida(id string, comida Comida) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	comida.FechaModificacion = time.Now().Format(time.RFC3339) // Actualiza la fecha de modificación

	update := bson.M{
		"$set": bson.M{
			"nombrecomida":      comida.NombreComida,
			"tipocomida":        comida.TipoComida,
			"descripcioncomida": comida.DescripcionComida,
			"imagenvideo":       comida.ImagenVideo,
			"activo":            comida.Activo,
			"fechamodificacion": comida.FechaModificacion,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("comida no encontrada")
	}

	return nil
}

// DeleteComida elimina una comida por su ID
func DeleteComida(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("comida no encontrada")
	}

	return nil
}
