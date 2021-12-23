package DB

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)
type APOD struct {
	gorm.Model
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}


const (
	dbConnStr    = "host=%s user=%s dbname=%s sslmode=disable password=%s port=%s"
)
var Db *gorm.DB //база данных



func Create(){
	// get env variables
	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort:=os.Getenv("db_port")
	// GORM
	// строка подключения
	dbUri := fmt.Sprintf(dbConnStr, dbHost, userName, "postgres", password,dbPort) //Создать строку подключения
	Db, err := gorm.Open("postgres", dbUri)
	if err!=nil {
		fmt.Println(err)
	}
	Db.Exec(fmt.Sprintf("CREATE DATABASE nasa;"))
	dbUri = fmt.Sprintf(dbConnStr, dbHost, userName, dbName, password,dbPort) //Создать строку подключения
	Db, err = gorm.Open("postgres", dbUri)
	Db.Debug().AutoMigrate(&APOD{}) //Миграция базы данных
}



func GetDB() *gorm.DB {
	return Db
}

func CheckError(err error){
	if err != nil {
		log.Println(err)
	}
}