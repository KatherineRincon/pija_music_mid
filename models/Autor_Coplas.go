package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
	

)

// Autor_Coplas representa la estructura de un usuario en la base de datos
type Autor_Cloplas struct {
	Id               uint      `gorm:"column:id_autor;primaryKey;autoIncrement" json:"id"`
	NombreAutor      string    `gorm:"column:nombre_autor" json:"nombre_autor"`
	ImagenVideo      string    `gorm:"column:imagen_video;type:text" json:"imagen_video"`
	BiografiaAutor   string    `gorm:"column:biografia_autor" json:"biografia_autor"`
	Activo           bool      `gorm:"column:activo" json:"activo"`
	FechaCreacion    time.Time `gorm:"column:fecha_creacion;autoCreateTime" json:"fecha_creacion"`
	FechaModificacion time.Time `gorm:"column:fecha_modificacion;autoUpdateTime" json:"fecha_modificacion"`
}

// Agregar un nuevo autor
func AddAutorCopla(db *gorm.DB, autor Autor_Cloplas) (uint, error) {
	result := db.Create(&autor)
	if result.Error != nil {
		return 0, result.Error
	}
	return autor.Id, nil
}

// Obtener un autor por ID
func GetAutorCopla(db *gorm.DB, id uint) (Autor_Cloplas, error) {
	var autor Autor_Cloplas
	result := db.First(&autor, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Autor_Cloplas{}, errors.New("autor no encontrado")
	}
	return autor, result.Error
}

// Obtener todos los autores
func GetAllAutorCoplas(db *gorm.DB) ([]Autor_Cloplas, error) {
	var autores []Autor_Cloplas
	result := db.Find(&autores)
	if result.Error != nil {
		return nil, result.Error
	}
	return autores, nil
}

// Actualizar un autor existente
func UpdateAutorCopla(db *gorm.DB, id uint, autorData Autor_Cloplas) error {
	var autor Autor_Cloplas
	if err := db.First(&autor, id).Error; err != nil {
		return errors.New("autor no encontrado")
	}
	autor.NombreAutor = autorData.NombreAutor
	return db.Save(&autor).Error
}

// Eliminar un autor por ID
func DeleteAutorCopla(db *gorm.DB, id uint) error {
	result := db.Delete(&Autor_Cloplas{}, id)
	if result.RowsAffected == 0 {
		return errors.New("no se encontró el autor para eliminar")
	}
	return result.Error
}