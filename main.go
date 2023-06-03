package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	Id int `gorm:"primaryKey autoIncrement" json:"id"`
	Name string `json:"name"`
	PhotoProfile string `json:"photoProfile"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func main() {
	connectDatabase()
	e := echo.New()
	e.GET("/users", GetUserController)
	e.POST("/users", InsertUserController)
	e.Start(":8000")
}

func InsertUserController(c echo.Context) error {
	var userInput User
	c.Bind(&userInput)

	result := DB.Create(&userInput)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Message: "success",
		Data: userInput,
	})
}

func GetUserController(c echo.Context) error {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data: users,
	})
}

func connectDatabase(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja4_twitter?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Error")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&User{})
}
