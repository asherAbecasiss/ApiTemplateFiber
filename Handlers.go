package main

import (
	"github.com/gofiber/fiber/v2"
)





func (s *ApiServer) GetInfo(c *fiber.Ctx) error {

	return c.JSON("info")
}
