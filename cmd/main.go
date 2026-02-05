package main

import (
	"fmt"
	"log"

	"github.com/chonlasit2000/rbac-hexagonalRBAC/internal/adapter/config"
	"github.com/chonlasit2000/rbac-hexagonalRBAC/internal/adapter/storage/postgres"
	"github.com/chonlasit2000/rbac-hexagonalRBAC/internal/adapter/storage/redis"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Println("Config loaded successfully")

	// Connect Database
	db, err := postgres.NewPostgresDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to Database" + db.Name())

	// Connect Redis
	rdb, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis" + rdb.Options().Addr)
}
