package main

import (
	"errors"
	"mae/database"
	"github.com/gofiber/fiber/v2"
	
)


func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!")
}


func Routes(app *fiber.App) {

	app.Post("/foodlist", AddProduct)
	app.Get("/foodlist", GetProductName)
	app.Get("/foodlist/:id", GetProduct)
	app.Delete("/foodlist/:id", Delete)
	app.Put("/foodlist/:id", Update)
	// app.Post("/user", AddUser)
	// app.Get("/user", GetUser)
	// app.Get("/user/:id", GetUsers)
	// app.Delete("/user/:id", Delete)
	// app.Put("/user/:id", Update)
	app.Get("/images", GetImages)
	app.Get("/images/:id", GetImage)
}


//foodlist
func AddProduct(c *fiber.Ctx) error {
	var food database.FoodList
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}

	database.DB.Create(&food)
	return c.JSON(&food)
}

func GetProductName(c *fiber.Ctx) error {
	var food []database.FoodList

	database.DB.Find(&food)
	return c.JSON(&food)
}

func FindProduct(id int, food *database.FoodList) error {
	database.DB.Find(&food, "id=?", id)
	if food.ID == 0 {
		return errors.New("ProductId does not existed")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var food database.FoodList
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(&food)

}

func Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food database.FoodList
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Delete(&food)
	return c.SendString("Deleted product")
}

func Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var food database.FoodList
	if err != nil {
		return c.SendString(err.Error())
	}
	if err := FindProduct(id, &food); err != nil {
		return c.SendString(err.Error())
	}
	if err := c.BodyParser(&food); err != nil {
		return c.SendString(err.Error())
	}
	database.DB.Save(&food)
	return c.JSON(&food)
}
//images
func GetImages(c *fiber.Ctx) error {
	 je := []database.Images{}
	database.DB.Find(&je)
	return c.JSON(je)
	}

func GetImage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	je := database.Images{}
	database.DB.Find(&je, "image_id", id)
	return c.JSON(je)

}

func main() {

	database.Migration()

	app := fiber.New()
	app.Get("/", Welcome)
	Routes(app)
	app.Listen(":3000")

}
