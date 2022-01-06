package promotion

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(cnx *gorm.DB) Repository {
	return Repository{db: cnx}
}

func (r Repository) GetAll() (*[]Promocion, error) {
	var data []Promocion
	result := r.db.Model(Promocion{}).Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}

	return &data, nil
}

func (r Repository) Create(promotion Promocion) (*Promocion, error) {
	result := r.db.Create(&promotion)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, errors.New("Error al agregar la promoción")
	}

	return &promotion, nil
}

func (r Repository) SearchDate(inicio, fin time.Time) (*[]Promocion, error) {
	var data []Promocion
	result := r.db.Debug().Model(Promocion{}).Where("fecha_inicio >= ? AND fecha_fin <= ?", inicio, fin).Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}

	return &data, nil
}
