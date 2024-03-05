package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shinjangwoon/go/go_fiber/database"
	"github.com/shinjangwoon/go/go_fiber/model"
)

// Blog list
func BlogList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog list",
	}
	time.Sleep(time.Millisecond * 300)
	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

// Blog details page
func BlogDetail(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"

		c.Status((404))
		return c.JSON(context)

	}

	context["record"] = record
	context["statusText"] = "OK"
	context["msg"] = "Blog Detail"
	c.Status(200)
	return c.JSON(context)
}

// Add a Blog into Database
func BlogCreate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "ADD Blog",
	}

	record := new(model.Blog)

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request: ")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	result := database.DBConn.Create(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved success."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Delete a Blog
func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Update Blog",
	}

	//http://localhost:8000/2

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("record not found")

		context["statusText"] = ""
		context["msg"] = "Record not found"
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error parsing")
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")
	}

	context["msg"] = "Record Updated Success!"
	context["data"] = record

	c.Status(200)
	return c.JSON(context)

}

func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not found")
		context["msg"] = "Record not found"

		return c.JSON(context)

	}

	result := database.DBConn.Delete(record)

	if result.Error != nil {
		context["msg"] = "Something went wrong"
		return c.JSON(context)

	}

	context["statusText"] = "OK"
	context["msg"] = "Record deleted success!"
	c.Status(200)
	return c.JSON(context)

}
