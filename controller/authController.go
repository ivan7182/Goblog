package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kingztech2019/blogbackend/database"
	"github.com/kingztech2019/blogbackend/models"
	"github.com/kingztech2019/blogbackend/util"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}
	var UserData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("unable to parse body")
	}
	//Check Password
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password must be greater than 6 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "invalid Email Address",
		})
	}

	// Check If email already exists in the database
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&UserData)
	if UserData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Already Exist",
		})
	}

	user := models.User{
		FirstName: strings.TrimSpace(data["first_name"].(string)),
		LastName:  strings.TrimSpace(data["last_name"].(string)),
		Phone:     strings.TrimSpace(data["phone"].(string)),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	err := user.SetPassword(data["password"].(string))
	if err != nil {
		log.Println(err)
	}
	err = database.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account Created Successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("tidak dapat mem-parse body")
		return err
	}
	var user models.User
	if err := database.DB.Where("email=?", data.Email).First(&user).Error; err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Alamat Email tidak ditemukan, silakan buat akun",
		})
	}

	if err := user.ComparePassword(data.Password); err != nil {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "kata sandi salah",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have successfully login",
		"user":    user,
	})

}

type claims struct {
	jwt.StandardClaims
}
