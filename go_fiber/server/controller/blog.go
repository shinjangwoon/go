package controller

import (
	"github.com/gofiber/fiber/v2"
)

// Blog list
func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog list",
	}

	c.Status(200)
	return c.JSON(context)
}

// Add a Blog into Database
func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "ADD Blog",
	}

	c.Status(201)
	return c.JSON(context)
}

// Delete a Blog
func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Update Blog",
	}

	c.Status(200)
	return c.JSON(context)

}

func BlogDelete(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Delete Blog for the given ID",
	}

	c.Status(200)
	return c.JSON(context)

}
