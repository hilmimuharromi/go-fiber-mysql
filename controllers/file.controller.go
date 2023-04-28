package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UploadFileHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("upload")
	if err != nil {
		return err
	}

	c.SaveFile(file, "public/upload/"+file.Filename)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"file": file.Filename}})

}
