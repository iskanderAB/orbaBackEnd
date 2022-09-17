package Config

import (
	"orba-back-end/entities"
	

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connection() {
	var err error
	DB , err = gorm.Open(postgres.Open("host=localhost user=postgres password=mohamed1998 dbname=ORBA_DB port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&entities.User{})
	DB.AutoMigrate(&entities.Additional{})
	DB.AutoMigrate(&entities.Category{})
	DB.AutoMigrate(&entities.Command{})
	DB.AutoMigrate(&entities.Command_line{})
	DB.AutoMigrate(&entities.Favorite{})
	DB.AutoMigrate(&entities.Location{})
	DB.AutoMigrate(&entities.Notification{})
	DB.AutoMigrate(&entities.Option{})
	DB.AutoMigrate(&entities.Product{})
	DB.AutoMigrate(&entities.Provider{})
	DB.AutoMigrate(&entities.Reclamation{})
	DB.AutoMigrate(&entities.Rote{})
	DB.AutoMigrate(&entities.Type{})
	

}