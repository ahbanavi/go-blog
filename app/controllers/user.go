package controllers

import (
	"github.com/ahbanavi/go-blog/app/models"
	"github.com/ahbanavi/go-blog/app/validator"
	db "github.com/ahbanavi/go-blog/database"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	db.DB.First(&user, id)
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := validator.Validate(user)
	if err != nil {
		return err
	}
	db.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	db.DB.Where("id = ?", id).Updates(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(models.User)
	db.DB.Where("id = ?", id).Delete(&user)
	return c.JSON(user)
}
