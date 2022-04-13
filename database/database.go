package database

import (
	"fmt"
	"log"
	
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var DNS = "host=localhost user=postgres password=admin dbname=fooddelivery sslmode=disable"

func Migration() {
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal("not connected to the database")
	}
	fmt.Print("connected to the database")
	DB.AutoMigrate(&FoodList{}, &Images{})
}

type FoodList struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stars       string `json:"stars"`
	Image         string `json:"Image"`
	Location    string `json:"location"`
	CreatedAt   string `json:"createdat"`
	UpdatedAt   string `json:"updatedat"`
	TypeId      int    `jason:"typeid"`
}
//images
type Images struct {
	Image_id uint `json:"image_id"`
	Image_name string ` json:"imag_name"`
	Imaege_url string `json:"image_url"`
}

// type User struct{
// 	ID uint `json:"id" gorm:"primaryKey"`
// 	Name string `json:"name"`
// 	Address string `json:"address"`
// }

func AddProduct(c *fiber.Ctx) error {
	var food FoodList
	if err := c.BodyParser(&food); err != nil {
		return err

	}
	DB.Create(&food)
	return c.JSON(&food)
}

//images
func AddImages(c *fiber.Ctx) error {
	var je Images
	if err := c.BodyParser(&je); err != nil {
		return err

	}
	DB.Create(&je)
	return c.JSON(&je) 
}

//user
// func AddUser(c *fiber.Ctx) error {
// 	var user User
// 	if err := c.BodyParser(&user); err != nil {
// 		return err

// 	}
// 	DB.Create(&user)
// 	return c.JSON(&user) 
// }

