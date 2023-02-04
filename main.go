package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	app := fiber.New()
	app.Get("/GetHead/:id", getHead)

	err := connectDB()
	if err != nil {
		log.Fatal("DB connection error:", err.Error())
	}
	err = migrateDB()
	if err != nil {
		log.Fatal("DB migration error", err.Error())
	}
	mongoClient.Database(os.Getenv("DB_DATABASE")).Collection("list").InsertOne(context.TODO(), bson.D{
		{Key: "_id", Value: primitive.NewObjectID()},
		{Key: "created_at", Value: time.Now()},
		{Key: "asdf", Value: "What ever hope it work"},
	})

	// log.Fatal(app.Listen(":3000"))
}
