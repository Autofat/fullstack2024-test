package handlers

import (
	"PeepL-Test/database"
	"PeepL-Test/models"
	"PeepL-Test/pkg/redis"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func CreateClient(c fiber.Ctx) error {
	client := new(models.My_client)
	if err := c.Bind().Body(client); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":"true",
			"message":"Invalid request body",
		})
	}
	if client.Name == ""{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":"true",
			"message":"Name is required",
		})
	}
	client.Slug = strings.ToLower(strings.ReplaceAll(client.Name, " ", "-")) // peepl-techologies-indonesia
	client.ClientPrefix = strings.ToUpper(client.Slug[:4]) //PEEP

	if client.IsProject == ""{
		client.IsProject = "0"
	}
	if client.SelfCapture == ""{
		client.SelfCapture = "1"
	}
	if client.ClientLogo == ""{
		client.ClientLogo = "no-image.jpg"
	}

	jsonData,_ := json.Marshal(client)
	err:=redis.RedisClient.Set(redis.Ctx, client.Slug, jsonData, 0).Err()
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":"true",
			"message":"Failed cache data",
		})
	}

	if err := database.DB.Create(&client).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":"true",
			"message":"Failed create client",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"error":"false",
		"message":"Client created successfully",
	})

}

func UpdateClient(c fiber.Ctx) error{
	id := c.Params("id")
	var client models.My_client
	if err := database.DB.First(&client, id).Error; err!= nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":"true",
			"message":"Client not found",
		})
	}
	oldSlug := client.Slug
	update := new(models.My_client)
	if err := c.Bind().Body(update); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":"true",
			"message":"Invalid request body",
			"data":err,
		}) 
	}

	if update.Name != ""{
		update.Slug = strings.ToLower(strings.ReplaceAll(update.Name, " ", "-"))
		update.ClientPrefix = strings.ToUpper(update.Slug[:4])
	}

	if err := database.DB.Model(&client).Updates(update).Error; err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":"true",
			"message":"Failed to update client",
			"data":err,
		})
	}

	redis.RedisClient.Del(redis.Ctx, oldSlug)
	jsonData, _ := json.Marshal(client)
	redis.RedisClient.Set(redis.Ctx, client.Slug, jsonData, 0)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"error":"false",
		"message":"Client updatd successfully",
	})
}

func DeleteClient(c fiber.Ctx) error{
	id := c.Params("id")
	var client models.My_client
	if err := database.DB.First(&client, id).Error; err!= nil{
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":"true",
			"message":"Client not found",
		})
	}

	if err := database.DB.Delete(&client).Error; err != nil{
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":"true",
			"message":"Failed to delete client",
		})
	}

	redis.RedisClient.Del(redis.Ctx, client.Slug)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"error":"false",
		"message":"Client successfuly deleted",
	})
}