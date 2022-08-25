package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	coll := client.Database("gomongodbreact").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: "Pablo",
	}})

	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "users from backend",
		})
	})

	app.Listen(":" + port)
	fmt.Println("Serving run in port 4000")
}
