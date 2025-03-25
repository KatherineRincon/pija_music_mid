package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

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

func AddArtista(artista *Artista) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(artista)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateArtista(artista *Artista) error {
	o := orm.NewOrm()
	if _, err := o.Update(artista); err != nil {
		return err
	}
	return nil
}

func GetArtistaById(id int) (*Artista, error) {
	o := orm.NewOrm()
	artista := Artista{Id_Artista: id}
	if err := o.Read(&artista); err != nil {
		return nil, err
	}
	return &artista, nil
}

func DeleteArtista(id int) error {
	o := orm.NewOrm()
	if _, err := o.Delete(&Artista{Id_Artista: id}); err != nil {
		return err
	}
	return nil
}
