package main

import (
	"akademia-api/confs"
	"akademia-api/db"
	"akademia-api/server"
	"log"
	"time"
)

func main() {

	err := confs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	// Initialize the database connection
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// initialize Redis
	redisServer := server.NewRedisServer(db)
	redisClient := redisServer.Start()
	defer redisServer.Close()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
			}
		}()
		time.Sleep(5 * time.Second)
		log.Println("Starting RabbitMQ server...")
		rabbitServer := server.NewRabbitMQServer(db, redisClient)
		rabbitServer.Start()
	}()

	// Initialize the server
	server := server.NewServer(db, redisClient)
	server.Start()

}
