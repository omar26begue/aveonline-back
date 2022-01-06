package invoice

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

func (r Repository) GetPayment(start, end time.Time) (*[]Payment, error) {
	var data []Payment
	result := r.db.Debug().Model(&Invoice{}).Select("to_char(fecha_creacion, 'YYYY/MM/DD') AS fecha, sum(pago_total) as total").Group("fecha").Where("fecha_creacion >= ?", start).Where("fecha_creacion <= ?", end).Find(&data)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, gorm.ErrRecordNotFound
	}
	return &data, nil
}
