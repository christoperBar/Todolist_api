package todocontroller

import (
	"net/http"

	"github.com/christoperBar/Todolist_api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AllTodoLists(c *fiber.Ctx) error {
	var todolists []models.Todo
	models.DB.Find(&todolists)

	return c.Status(fiber.StatusOK).JSON(todolists)
}

func GetTodoList(c *fiber.Ctx) error {
	id := c.Params("id")
	var todolist models.Todo
	if err := models.DB.First(&todolist, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data not Found",
			})
		}
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todolist)
}

func AddTodoList(c *fiber.Ctx) error {

	var todolist models.Todo
	if err := c.BodyParser(&todolist); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&todolist).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todolist)
}

func UpdateTodoList(c *fiber.Ctx) error {

	id := c.Params("id")

	var todolist models.Todo
	if err := c.BodyParser(&todolist); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id = ?", id).Updates(&todolist).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tidak berhasil mengupdate data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func DeleteTodoList(c *fiber.Ctx) error {

	id := c.Params("id")

	var todolist models.Todo
	if models.DB.Delete(&todolist, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}
