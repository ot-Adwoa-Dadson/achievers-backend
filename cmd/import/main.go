package main

import (
	//"context"
	//"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"


	//"github.com/gocarina/gocsv"
	"fellowship-backend/internal/config"
	//"fellowship-backend/internal/importer"
	"fellowship-backend/internal/jobs"
	"fellowship-backend/internal/routes"
)

func main() {
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

	routes.RegisterRoutes(router, memberCol, userCol)

	// Start birthday cron
	jobs.StartBirthdayCron(memberCol, notificationCol)

	router.Run(":8080")

}
