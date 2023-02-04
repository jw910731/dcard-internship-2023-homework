package main

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/joho/godotenv/autoload"
)

var mongoClient *mongo.Client

func connectDB() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(os.Getenv("DB_URI"))
	mongoClient, err = mongo.Connect(ctx, clientOpts)
	return
}

func migrateDB() error {
	db := mongoClient.Database(os.Getenv("DB_DATABASE"))
	list_coll := db.Collection("list")
	ttl_idx_opt := mongo.IndexModel{
		Keys:    bson.D{{Key: "created_at", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(60 * 60 * 24 * 1),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := list_coll.Indexes().CreateOne(ctx, ttl_idx_opt)
	return err
}
