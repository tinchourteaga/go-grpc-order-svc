package db

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector struct {
	DB *gorm.DB
}

func NewDatabaseConnection() Connector {
	db, err := gorm.Open(postgres.Open(viper.GetString("DB_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("cannot connect to database")
	}

	//db.AutoMigrate(&models.Product{}, &models.StockDecreaseLog{})

	return Connector{DB: db}
}
