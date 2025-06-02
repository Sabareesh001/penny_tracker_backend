package database

import (
	"fmt"
	"os"
	"strconv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect()  *gorm.DB{


	USERNAME := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	SOCKET_NUMBER := os.Getenv("SOCKET_NUMBER")
	DATABASE := os.Getenv("DATABASE")
	CHARSET := os.Getenv("CHARSET")
	PARSE_TIME := os.Getenv("PARSE_TIME")
	LOC:=os.Getenv("LOC")

	dsn := USERNAME+":"+PASSWORD+"@"+PROTOCOL+"("+SOCKET_NUMBER+")"+"/"+DATABASE+"?"+"charset="+CHARSET+"&parseTime="+PARSE_TIME+"&loc="+LOC

	db,errGormDB := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if errGormDB!=nil {
		fmt.Println("Cannot connect to the Database 🥹😖")
		panic(errGormDB)
	}

	fmt.Println("Connected to the Database 🎉")
	
	fmt.Println("Configurating Database Connection...")
	sqlDB,errSqlDB := db.DB();
	if errSqlDB!=nil{
		fmt.Println("Cannot connect to the Database 🥹😖")
		panic(errSqlDB)
	}
 
    MAX_IDLE_CONNECTIONS,maxIdleConnError := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTIONS"))
    MAX_OPEN_CONNECTIONS,maxOpenConnError := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTIONS"))

    if maxIdleConnError!=nil{
		fmt.Println("Invalid format for MAX_IDLE_CONNECTIONS")
		panic(maxIdleConnError)
	}

	if maxOpenConnError!=nil{
		fmt.Println("Invalid format for MAX_OPEN_CONNECTIONS")
		panic(maxOpenConnError)
	}

	sqlDB.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
    sqlDB.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
	
	return db;
}