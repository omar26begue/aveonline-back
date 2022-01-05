package invoice

import (
	"aveonline-back/modules/medicine"
	"aveonline-back/modules/promotion"
	"time"
)

type Invoice struct {
	Id            uint                `json:"id" gorm:"primaryKey;auto_increment"`
	FechaCreacion time.Time           `json:"fecha_creacion" gorm:"not null"`
	PagoTotal     float32             `json:"pago_total" gorm:"not null"`
	IdPromotion   uint                `json:"id_promotion" gorm:"not null"`
	Promotion     promotion.Promocion `gorm:"foreignKey:id_promotion"`
	Medicines     []medicine.Medicine `json:"medicines" gorm:"many2many:invoice_medicines"`
}
