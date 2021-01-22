package dblayer

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname string) (*DBORM, error) {
	db, err := gorm.Open(mysql.Open(dbname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("DB Connecting ...")
	return &DBORM{
		DB: db,
	}, err
}