package controller

import "github.com/gofiber/fiber/v2"

type IController interface {
	GetAll(c *fiber.Ctx) error
	GetOne(c *fiber.Ctx) error
	Add(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	MountRouter(c fiber.Router)
}
