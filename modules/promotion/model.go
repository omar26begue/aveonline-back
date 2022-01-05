package promotion

import "time"

type Promocion struct {
	Id          uint      `json:"id" gorm:"primaryKey;auto_increment"`
	Descripcion string    `json:"descripcion" gorm:"size:100; not null"`
	Porcentaje  float32   `json:"porcentaje" gorm:"not null"`
	FechaInicio time.Time `json:"fecha_inicio" gorm:"not null"`
	FechaFin    time.Time `json:"fecha_fin" gorm:"not null"`
}

type PromocionRequest struct {
	Descripcion string  `json:"descripcion" validate:"required" example:"Prueba"`
	Porcentaje  float32 `json:"porcentaje" validate:"required,gt=0,lte=70" example:"1.5"`
	FechaInicio string  `json:"fecha_inicio" validate:"required" example:"2021-12-05"`
	FechaFin    string  `json:"fecha_fin" validate:"required" example:"2021-12-07"`
}
