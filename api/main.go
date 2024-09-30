package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ogbonnaEmmanuel/GkeDeploy/api/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func main()  {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress, // Get Redis address from config or env
	})

	pong, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Printf("Redis ping response: %s", pong)

	log.Println(config.RedisAddress)
	
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MONGO_URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Check the connection
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("MongoDB is not reachable: %v", err)
	}

	db_collections := []string{"unicollection", "board_access_collection", "users"}
	err = util.CheckAndCreateDatabase(mongoClient, "uni", db_collections)
	if err != nil {
		log.Fatalf("Error occured creating collections: %v", err)
	}

	r := gin.Default()

	// Define the /hello endpoint
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
