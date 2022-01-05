package config

import (
	"aveonline-back/modules/invoice"
	"aveonline-back/modules/medicine"
	"aveonline-back/modules/promotion"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDatabase() (*invoice.Repository, *promotion.Repository, *medicine.Repository) {
	dsn := "host="+viper.GetString("DB_HOST")+" user="+viper.GetString("DB_USER")+" password="+viper.GetString("DB_PASSWORD")+" dbname="+viper.GetString("DB_NAME")+" port="+viper.GetString("DB_PORT")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}

	migration := NewMigration(db)
	migration.InitMigration()

	repoInvoice := invoice.NewRepository(db)
	repoPromotion := promotion.NewRepository(db)
	repoMedicine := medicine.NewRepository(db)

	return &repoInvoice, &repoPromotion, &repoMedicine
}
