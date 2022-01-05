package medicine

type Medicine struct {
	Id       uint    `json:"id" gorm:"primaryKey;auto_increment"`
	Name     string  `json:"name" gorm:"size:50; not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Location string  `json:"location" gorm:"size:50; not null"`
}

type MedicineRequest struct {
	Name     string  `json:"name" validate:"required"`
	Price    float64 `json:"price" validate:"required,gt=0"`
	Location string  `json:"location" validate:"required"`
}
