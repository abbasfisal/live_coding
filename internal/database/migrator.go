package database

import (
	"fmt"
	"live_coding/entity"
	"log"
)

func Migrate() {
	Connect()
	db := Get()
	db.Migrator().DropTable(&entity.Address{}, &entity.User{})
	fmt.Println("[Drop Table ] Success")

	err := db.AutoMigrate(&entity.User{}, &entity.Address{})
	if err != nil {
		log.Fatal("[Migrate] failed :", err, "\n")
		return
	}
	fmt.Printf("[Migrate] success \n")
}
