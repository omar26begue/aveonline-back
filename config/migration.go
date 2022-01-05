package config

import (
	"aveonline-back/modules/invoice"
	"aveonline-back/modules/medicine"
	"aveonline-back/modules/promotion"
	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func NewMigration(cnx *gorm.DB) *Migration {
	return &Migration{db: cnx}
}

func (m *Migration) InitMigration() {
	m.db.AutoMigrate(&medicine.Medicine{}, &promotion.Promocion{}, &invoice.Invoice{})
}
