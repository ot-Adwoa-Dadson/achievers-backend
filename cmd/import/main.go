package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"fellowship-backend/internal/config"
	"fellowship-backend/internal/jobs"
	"fellowship-backend/internal/routes"
)

type MemberCSV struct {
	FullName               string `csv:"fullName"`
	Phone                  string `csv:"phone"`
	Email                  string `csv:"email"`
	HomeAddress            string `csv:"homeAddress"`
	DateOfBirth            string `csv:"dateOfBirth"`
	Occupation             string `csv:"occupation"`
	CurrentEmployer        string `csv:"currentEmployer"`
	GuardianName           string `csv:"guardianName"`
	GuardianRelationship   string `csv:"guardianRelationship"`
	GuardianPhone          string `csv:"guardianPhone"`
	GuardianAltPhone       string `csv:"guardianAlternativePhone"`
	GuardianEmail          string `csv:"guardianEmail"`
	SeniorCell             string `csv:"seniorCell"`
	FoundationSchoolStatus string `csv:"foundationSchoolStatus"`
	LeadershipRole         string `csv:"leadershipRole"`
	DesignationCell        string `csv:"designationCell"`
}
type Guardian struct {
	Name             string `bson:"name"`
	Relationship     string `bson:"relationship"`
	Phone            string `bson:"phone"`
	AlternativePhone string `bson:"alternativePhone"`
	Email            string `bson:"email"`
}

type FellowshipInfo struct {
	SeniorCell             string `bson:"seniorCell"`
	FoundationSchoolStatus string `bson:"foundationSchoolStatus"`
	LeadershipRole         string `bson:"leadershipRole"`
	DesignationCell        string `bson:"designationCell"`
}
type Member struct {
	FullName        string        `bson:"fullName"`
	Phone           string        `bson:"phone"`
	Email           string        `bson:"email"`
	HomeAddress     string        `bson:"homeAddress"`
	DateOfBirth     time.Time     `bson:"dateOfBirth"`
	Occupation      string        `bson:"occupation"`
	CurrentEmployer string        `bson:"currentEmployer"`
	Guardian        Guardian      `bson:"guardian"`
	Fellowship      FellowshipInfo `bson:"fellowship"`
	CreatedAt       time.Time     `bson:"createdAt"`
	UpdatedAt       time.Time     `bson:"updatedAt"`
}

// parseDOB parses MM/DD/YYYY
func parseDOB(dob string) (time.Time, error) {
	dob = strings.TrimSpace(dob)
	if dob == "" {
		return time.Time{}, fmt.Errorf("empty DOB")
	}
	t, err := time.Parse("1/2/2006", dob)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid DOB: %s", dob)
	}
	return t, nil
}

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





// package main

// import (
// 	"context"
// 	"encoding/csv"
// 	"log"
// 	"os"
// 	"strings"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"

// 	"fellowship-backend/internal/config"
// 	"fellowship-backend/internal/models"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// Load .env
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("⚠️ No .env file found, using system environment variables")
// 	}

// 	mongoURI := os.Getenv("MONGO_URI")
// 	if mongoURI == "" {
// 		log.Fatal("MONGO_URI not set")
// 	}

// 	client, err := config.ConnectMongo(mongoURI)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db := client.Database("achievers")
// 	memberCol := db.Collection("members")

// 	// ---- 1️⃣ Wipe existing members ----
// 	if _, err := memberCol.DeleteMany(context.TODO(), bson.M{}); err != nil {
// 		log.Fatal("Failed to clear members collection:", err)
// 	}
// 	log.Println("✅ Cleared existing members")

// 	// ---- 2️⃣ Open CSV ----
// 	file, err := os.Open("data/members.csv")
// 	if err != nil {
// 		log.Fatal("Failed to open CSV:", err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		log.Fatal("Failed to read CSV:", err)
// 	}

// 	if len(records) < 2 {
// 		log.Fatal("CSV is empty or missing header")
// 	}

// 	header := records[0]
// 	data := records[1:]

// 	log.Printf("✅ Found %d rows to import\n", len(data))

// 	// ---- 3️⃣ Process and insert each row ----
// 	for i, row := range data {
// 		rowMap := map[string]string{}
// 		for j, col := range header {
// 			rowMap[col] = row[j]
// 		}

// 		// Parse dateOfBirth (MM/DD/YYYY)
// 		var dob time.Time
// 		if raw := strings.TrimSpace(rowMap["dateOfBirth"]); raw != "" {
// 			dob, err = time.Parse("1/2/2006", raw) // Go format: 1/2/2006 = MM/DD/YYYY
// 			if err != nil {
// 				log.Printf("Row %d: failed to parse DOB '%s': %v\n", i+2, raw, err)
// 				dob = time.Time{}
// 			}
// 		}

// 		member := models.Member{
// 			FullName:    strings.TrimSpace(rowMap["fullName"]),
// 			Phone:       strings.TrimSpace(rowMap["phone"]),
// 			Email:       strings.TrimSpace(rowMap["email"]),
// 			HomeAddress: strings.TrimSpace(rowMap["homeAddress"]),
// 			DateOfBirth: dob,
// 			Occupation:  strings.TrimSpace(rowMap["occupation"]),
// 			CurrentEmployer: strings.TrimSpace(rowMap["currentEmployer"]),
// 			Guardian: models.Guardian{
// 				Name:         strings.TrimSpace(rowMap["guardianName"]),
// 				Relationship: strings.TrimSpace(rowMap["guardianRelationship"]),
// 				Phone:        strings.TrimSpace(rowMap["guardianPhone"]),
// 				AlternativePhone: strings.TrimSpace(rowMap["guardianAlternativePhone"]),
// 				Email:        strings.TrimSpace(rowMap["guardianEmail"]),
// 			},
// 			Fellowship: models.FellowshipInfo{
// 				SeniorCell:             strings.TrimSpace(rowMap["seniorCell"]),
// 				FoundationSchoolStatus: strings.TrimSpace(rowMap["foundationSchoolStatus"]),
// 				LeadershipRole:         strings.TrimSpace(rowMap["leadershipRole"]),
// 				DesignationCell:        strings.TrimSpace(rowMap["designationCell"]),
// 			},
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 		}

// 		_, err := memberCol.InsertOne(context.TODO(), member)
// 		if err != nil {
// 			log.Printf("Row %d: failed to insert: %v\n", i+2, err)
// 		}
// 	}

// 	log.Println("✅ Import complete!")
// }
