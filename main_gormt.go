package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wonli/gormt"
	"github.com/wonli/gormt/config"
)

func main3() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("user")
	pass := os.Getenv("pass")
	dbname := os.Getenv("dbname")

	dbInfo := config.DBInfo{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: user,
		Password: pass,
		Database: dbname,
		Type:     0,
	}

	conf := config.Config{
		DBInfo:           dbInfo,
		PkgName:          "schema",
		OutDir:           "./models",
		DbTag:            "gorm",
		IsJsonTag:        true,
		IsNullToSqlNull:  true,
		TablePrefix:      "q_",
		StripTablePrefix: true,
		OutFileName:      "schema",
	}

	gormt.ExecuteConfig(&conf)
}
