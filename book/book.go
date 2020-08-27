package book

import (
	"../database"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Writer string `json:"writer"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var book []Book
	db.Find(&book)
	c.JSON(book)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	if book.Id == 0 {
		c.Status(500).Send("No Book Found with Id: ", id)
		return
	}
	db.Delete(book, id)
	c.Send("Book Successfully deleted.")
}
