package main

import (
	//"context"
	"log"
	"os"
	"fmt"
    "net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

     "github.com/gin-contrib/cors"
     "time"

	//"github.com/gocarina/gocsv"
	"fellowship-backend/internal/config"
	//"fellowship-backend/internal/importer"
	"fellowship-backend/internal/jobs"
	"fellowship-backend/internal/routes"
)

func main() {
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default for local development
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Fellowship backend running!")
    })

    log.Printf("Server starting on port %s...\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
	// LOAD .env FILE
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	client, err := config.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	// collection := client.Database("achievers").Collection("members")

	// file, err := os.Open("data/members.csv")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// var records []importer.MemberCSV
	// if err := gocsv.UnmarshalFile(file, &records); err != nil {
	// 	log.Fatal(err)
	// }

	// for i, record := range records {
	// 	member := importer.CsvToMember(record)

	// 	_, err := collection.InsertOne(context.TODO(), member)
	// 	if err != nil {
	// 		log.Printf("Row %d failed: %v\n", i+1, err)
	// 	}
	// }

	// fmt.Println("Members imported successfully")
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println("✅ Test email sent successfully!")
	  

	db := client.Database("achievers")

	memberCol := db.Collection("members")
	userCol := db.Collection("users")
	notificationCol := db.Collection("notifications")

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	routes.RegisterRoutes(router, memberCol, userCol)

	// Start birthday cron
	jobs.StartBirthdayCron(memberCol, notificationCol)

	// 🟢 Allow CORS for frontend development


	router.Run(":8080")

}
