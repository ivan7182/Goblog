package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kingztech2019/blogbackend/database"
	"github.com/kingztech2019/blogbackend/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("unable to parse body")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulation, your post is live ",
	})

}
