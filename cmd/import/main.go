package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fellowship-backend/internal/config"
	"fellowship-backend/internal/jobs"
	"fellowship-backend/internal/routes"
)

func main() {
	// Load .env if exists
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	// Get PORT from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get Mongo URI from env
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set in environment")
	}

	// Connect to MongoDB
	client, err := config.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	log.Println("✅ Connected to MongoDB")

	// Select database and collections
	db := client.Database("achievers")
	memberCol := db.Collection("members")
	userCol := db.Collection("users")
	notificationCol := db.Collection("notifications")

	

	// Initialize Gin router
	router := gin.Default()

	// CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:5173",
			"https://achievers-backend-eti6.onrender.com",
			"https://achievers-pcf-admin.netlify.app",
		}, // add production frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register all routes
	routes.RegisterRoutes(router, memberCol, userCol)

	// Start cron jobs
	jobs.StartBirthdayCron(memberCol, notificationCol)

	// Start server
	log.Printf("🚀 Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
