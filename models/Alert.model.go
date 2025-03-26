package models

import (
	"time"

)

type Alert struct {
	Type string
	Code string
	Body interface{}
}

// Artista representa la tabla Artista en la base de datos
type Artista struct {
	Id_Artista       int       `orm:"column(id);pk;auto"`           // ID del artista
	Nombre_Artistico string    `orm:"column(nombre_artistico)"`     // Nombre artístico
	Nombre_Real      string    `orm:"column(nombre_real)"`          // Nombre real
	Genero_Musical   string    `orm:"column(genero_musical)"`       // Género musical (lo cambié a string en vez de int)
	Biografia        string    `orm:"column(biografia)"`            // Biografía del artista
	Imagen_Video     string    `orm:"column(imagen_video);null"`    // Imagen o video (opcional)
	Redes_Sociales   string    `orm:"column(redes_sociales);null"`  // Redes sociales (opcional)
	Activo           bool      `orm:"column(activo);default(true)"` // Estado activo/inactivo del artista
	CreatedAt        time.Time `orm:"auto_now_add;type(datetime)"`  // Fecha de creación
	UpdatedAt        time.Time `orm:"auto_now;type(datetime)"`      // Fecha de actualización
}

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}


