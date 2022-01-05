package invoice

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

func (r Repository) GetAll() (*[]Invoice, error) {
	var data []Invoice
	result := r.db.Model(Invoice{}).Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}

	return &data, nil
}

func (r Repository) Create(invoice Invoice) (*Invoice, error) {
	result := r.db.Create(&invoice)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, errors.New("Error al agregar la factura")
	}

	return &invoice, nil
}
