package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pciruelos/go-react-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodbreact"))

	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "dale users from backend",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)

		coll := client.Database("gomongodbreact").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})
		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": "guardando usuario from post",
		})
	})

	app.Listen(":" + port)
	fmt.Println("Serving run in port 4000")
}
