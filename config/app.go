package config

import(
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)
var (
	db * gorm.DB
)
func Connect(){
	er := godotenv.Load()
	if er !=nil{
		panic("Fail to load .env file")
	}
	DB_HOST:=os.Getenv("DB_HOST")
	DB_USER:=os.Getenv("DB_USER")
	DB_PASS:=os.Getenv("DB_PASS")
	DB_NAME:=os.Getenv("DB_NAME")

	conn,err := gorm.Open(postgres.Open("host="+DB_HOST+" user="+DB_USER+" password="+DB_PASS+" dbname="+DB_NAME))
	if err != nil {
		panic(err)
	}
	db = conn
}
func GetDB() *gorm.DB{
	return db
}
func Disconnect() {
	dbConn,err:=db.DB()
	if err != nil{
		panic(err)
	}
	dbConn.Close()
}