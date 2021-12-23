package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"io"
	"log"
	"main.go/DB"
	"main.go/logger"
	"net/http"
	"os"
	"time"
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
	key_api string= "m4xVN4JznEeOlxdXdqZFtDGk4Gl7tX0s0RsoxaZS"
)
var db *gorm.DB //база данных
const (
	dbConnStr    = "host=%s user=%s dbname=%s sslmode=disable password=%s port=%s"
)

func init(){
	// logger
	logger.InitLocalLogger()
	logger.LocalLog.Println("start app")
}

func main() {
	DB.Create()
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	userName := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort:=os.Getenv("db_port")

	// GORM
	// строка подключения к DB
	dbUri := fmt.Sprintf(dbConnStr, dbHost, userName, dbName, password,dbPort)

	db, err = gorm.Open("postgres", dbUri)
	if err!=nil {
		fmt.Println(err)
	}
	defer db.Close()

	dt:=time.Now()
	path := "img/" + (dt.Format("20060102_150405")) + ".jpg"
	fmt.Println(path)
	logger.LocalLog.Println("Image will be download in :", path)



 	//ticker := time.NewTicker(time.Second)

	err= db.Create(&APOD{Copyright: "hello",
		Date:"hello" ,
		Explanation :"hello"  ,
		Hdurl        :"hello" ,
		MediaType    :"hello",
		ServiceVersion :"hello",
		Title          :"hello",
		URL            :"hello"}).Error
	var s = APOD{Copyright: "1",
		Date:"1" ,
		Explanation :"1"  ,
		Hdurl        :"1" ,
		MediaType    :"1",
		ServiceVersion :"1",
		Title          :"1",
		URL            :"1",
	}
	fmt.Println(s)
	err= db.Create(&s).Error


	//for {
	//	select {
	//	case t := <-ticker.C:

			// каждые 24 часа
			//if t.Hour()%24 == 0 {
			// для простоты делаем каждую минуту
	//		if t.Second()%60 == 0 {

				getPicture(key_api)

	//		}
	//	}
	//}
}


func getPicture(key string) (jsonResp APOD,err error){
	resp,err:=http.Get(fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s",key))
	defer resp.Body.Close()
	if err!=nil{
		logger.LocalLog.Println("No connect to API")
		return
	}
	//fmt.Println(resp.Body)
	err= json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err!=nil{
		logger.LocalLog.Println("Not correct format API")
		return
	}
	fmt.Println(&jsonResp)
	s:=jsonResp
	fmt.Println(s)


	err= db.Create(&s).Error

		if (len(jsonResp.URL)!= 0) && (jsonResp.MediaType=="image") {
			go downloadPic(jsonResp.URL)
		} else {
			logger.LocalLog.Println("no image today")
			}




	return
}


func downloadPic (link string)(err error){
	resp, err := http.Get(link)
	CheckError(err)
	defer resp.Body.Close()
	dt:=time.Now()
	path := "img/" + (dt.Format("20060102_150405")) + ".jpg"
	out, err := os.Create(path)
	CheckError(err)
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.LocalLog.Println(err)
	} else {
		logger.LocalLog.Println("downloaded image from URL", link)
	}
	return
}


//func getPictureName(link string)(string, err){
//	ss := strings.Split(link, "/")
//	fmt.Println(ss)
//
//}

func CheckError(err error){
	if err != nil {
		log.Println(err)
	}
}



