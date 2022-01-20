package database

import (
	"fmt"
	"strings"

	"github.com/juliazuin/AluraChallenge/app/model"
	config "github.com/juliazuin/AluraChallenge/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var c = config.NewConfig()

var username = strings.TrimSpace(c.Viper.GetString("database.username")) // os.Getenv("DB_USERNAME")
var password = c.Viper.GetString("database.password")                    // os.Getenv("DB_PASSWORD")
var host = c.Viper.GetString("database.host")                            // os.Getenv("DB_HOST")
var name = c.Viper.GetString("database.name")                            // os.Getenv("DB_NAME")
var port = c.Viper.GetString("database.port")                            // os.Getenv("DB_PORT")

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")/" + name + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	db.AutoMigrate(model.Receitas{}, model.Despesas{})

	return db
}
