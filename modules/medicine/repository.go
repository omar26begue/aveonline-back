package medicine

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(cnx *gorm.DB) Repository {
	return Repository{db: cnx}
}

func (r Repository) GetAll() (*[]Medicine, error) {
	var data []Medicine
	result := r.db.Model(Medicine{}).Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}

	return &data, nil
}

func (r Repository) Create(medicine Medicine) (*Medicine, error) {
	result := r.db.Create(&medicine)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, errors.New("Error al agregar la factura")
	}

	return &medicine, nil
}
