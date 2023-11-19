package postgres

import (
	"errors"
	"finalProject3/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbName   = "db_finalproject_3"
	db       *gorm.DB
	err      error
)

func GetDBConfig() gorm.Dialector {

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		port,
		user,
		password,
		dbName,
	)

	return postgres.Open(dbConfig)
}

func GetDBInstance() *gorm.DB {
	return db
}

func seedAdmin() {
	admin := &entity.User{
		FullName: "admin",
		Email:    "admin@gmail.com",
		Password: "admin12345",
		Role:     "admin",
	}
	err := admin.HashPassword()

	if err != nil {
		fmt.Println(err)
	}

	err2 := db.Create(admin).Error

	if err2 != nil {
		fmt.Println(err2)
	}

	log.Println("Admin account seed success!")
}

func init() {
	db, err = gorm.Open(GetDBConfig(), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Category{}, &entity.Task{})

	if db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			seedAdmin()
		}
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	log.Println("Connected to DB!")
}
