package main

import (
	"orba-back-end/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// //loding env varibales
	// dialect := os.Getenv("DIALECT")
	// host := os.Getenv("HOST")
	// dbport := os.Getenv("DBPORT")
	// user :=os.Getenv("USER")
	// dbname := os.Getenv("NAME")
	// password := os.Getenv("PASSWORD")

	// //database connection string
	// dburi := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s  ",host,user,dbname,password,dbport)
	// //connection to the database
	// db ,err = gorm.Open(dialect,dburi)
	// if err != nil{
	// 	log.Fatal(err)
	// } else{
	// 	fmt.Println("database opened seccessfully")
	// }
	// //colse the database connection
	// db.Close()
	// //make migration
	// db.AutoMigrate(&entities.User{})
	dsn := "host=localhost user=postgres password=mohamed1998 dbname=ORBA_DB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		db.AutoMigrate(&entities.User{})
	    db.AutoMigrate(&entities.Additional{})

		db.AutoMigrate(&entities.Category{})

		db.AutoMigrate(&entities.Command{})
		db.AutoMigrate(&entities.Command_line{})
		db.AutoMigrate(&entities.Favorite{})

		db.AutoMigrate(&entities.Location{})
		db.AutoMigrate(&entities.Notification{})
		db.AutoMigrate(&entities.Options{})
		db.AutoMigrate(&entities.Product{})
		db.AutoMigrate(&entities.Provider{})
		db.AutoMigrate(&entities.Reclamation{})
		db.AutoMigrate(&entities.Rote{})
		db.AutoMigrate(&entities.Type{})



				



	}
}
