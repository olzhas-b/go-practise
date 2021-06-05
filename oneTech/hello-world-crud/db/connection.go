package db

import (
	"bytes"
	"gorm.io/gorm"
	"log"
	"os/exec"
)
const (
	HOST="localhost"
	DBPORT="5432"
	USER="myuser"
	NAME="postgres"
	PASSWORD="user"
)
var db *gorm.DB
var err error
func InitDataBase() *gorm.DB {
	cmd := exec.Command("createdb", "-p", "5432", "-h", "127.0.0.1", "-U", "superuser", "-e", "test_db")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Output: %q\n", out.String())


	return db
	//config := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s TimeZone=Asia/Shanghai", HOST, USER, NAME, PASSWORD, DBPORT)
	//dataBase, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	//dataBase = dataBase.Exec("CREATE DATABASE test_db;")
	//
	//if err != nil {
	//	log.Println("failed open database")
	//	log.Println(err)
	//}
	//dataBase.AutoMigrate(&models.User{})
	////dataBase.Create(&models.User{
	////		ID: 1,
	////		Email:"oljas@gmail.com",
	////		Password: "oljas",
	////		Phone: "+7777777777",
	////		Name: "olzhas",
	////	})
	//var user models.User
	//dataBase.First(&user, 1)
	//fmt.Println(user)
	//return dataBase
}
func GetDataBase() *gorm.DB {
	if db == nil {
		db = InitDataBase()
	}
	return db
}

func capture(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}