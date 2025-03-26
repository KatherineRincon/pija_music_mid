package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
)

// Artista representa la estructura de un artista en la base de datos.
type Artista struct {
	Id_Artista       int       `orm:"column(id);pk;auto"`           // ID del artista (autoincremental)
	Nombre_Artistico string    `orm:"column(nombre_artistico)"`     // Nombre artístico
	Nombre_Real      string    `orm:"column(nombre_real)"`          // Nombre real
	Genero_Musical   string    `orm:"column(genero_musical)"`       // Género musical
	Biografia        string    `orm:"column(biografia)"`            // Biografía
	Imagen_Video     string    `orm:"column(imagen_video);null"`    // Imagen o video (opcional)
	Redes_Sociales   string    `orm:"column(redes_sociales);null"`  // Redes sociales (opcional)
	Activo           bool      `orm:"column(activo);default(true)"` // Estado activo/inactivo
	CreatedAt        time.Time `orm:"auto_now_add;type(datetime)"`  // Fecha de creación
	UpdatedAt        time.Time `orm:"auto_now;type(datetime)"`      // Fecha de actualización
}

// Registra el modelo en Beego ORM
func init() {
	orm.RegisterModel(new(Artista))
}

// AddArtista agrega un nuevo artista y retorna su ID.
func AddArtista(artista *Artista) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(artista)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetArtista obtiene un artista por su ID.
func GetArtista(id int) (*Artista, error) {
	o := orm.NewOrm()
	artista := Artista{Id_Artista: id}
	err := o.Read(&artista)
	if err == orm.ErrNoRows {
		return nil, errors.New("artista no encontrado")
	}
	return &artista, err
}

// GetAllArtistas obtiene todos los artistas.
func GetAllArtistas() ([]Artista, error) {
	o := orm.NewOrm()
	var artistas []Artista
	_, err := o.QueryTable(new(Artista)).All(&artistas)
	return artistas, err
}

// UpdateArtista actualiza la información de un artista.
func UpdateArtista(id int, artista *Artista) error {
	o := orm.NewOrm()
	artista.Id_Artista = id
	_, err := o.Update(artista)
	if err != nil {
		return errors.New("error al actualizar el artista")
	}
	return nil
}

// DeleteArtista elimina un artista por su ID.
func DeleteArtista(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Artista{Id_Artista: id})
	if err != nil {
		return errors.New("error al eliminar el artista")
	}
	return nil
}

func Login(username, password string) bool {
	for _, user := range usuarios { // Recorre la lista de usuarios
		if user.Nombre == username && user.Password == password {
			return true // Usuario encontrado
		}
	}
	return false // Usuario no encontrado
}

