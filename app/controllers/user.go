package controllers

import (
	"github.com/ahbanavi/go-blog/app/models"
	"github.com/ahbanavi/go-blog/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(users)
}
