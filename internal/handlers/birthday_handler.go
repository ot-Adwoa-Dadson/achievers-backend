// package handlers

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type BirthdayHandler struct {
// 	Collection *mongo.Collection
// }

// // GetUpcomingBirthdays returns members with birthdays in the next 7 days
// // func (h *BirthdayHandler) GetUpcomingBirthdays(c *gin.Context) {
// // 	now := time.Now()
// // 	day7 := now.AddDate(0, 0, 7)

// // 	// Filter by birthday (ignore year)
// // 	filter := bson.M{
// // 		"dateOfBirth": bson.M{
// // 			"$gte": now,
// // 			"$lte": day7,
// // 		},
// // 	}

// // 	cursor, err := h.Collection.Find(context.TODO(), filter)
// // 	if err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch birthdays"})
// // 		return
// // 	}
// // 	var members []map[string]interface{}
// // 	if err := cursor.All(context.TODO(), &members); err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse birthdays"})
// // 		return
// // 	}

// // 	c.JSON(http.StatusOK, members)
// // }

// func (h *BirthdayHandler) GetUpcomingBirthdays(c *gin.Context) {
// 	now := time.Now()
// 	currentMonth := now.Month()

// 	cursor, err := h.Collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch members"})
// 		return
// 	}

// 	var members []map[string]interface{}
// 	if err := cursor.All(context.TODO(), &members); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse members"})
// 		return
// 	}

// 	// Filter birthdays in the current month
// 	var upcoming []map[string]interface{}
// 	for _, m := range members {
// 		if dobVal, ok := m["dateOfBirth"].(time.Time); ok {
// 			if dobVal.Month() == currentMonth {
// 				upcoming = append(upcoming, m)
// 			}
// 		} else if dobStr, ok := m["dateOfBirth"].(string); ok {
// 			// in case date is stored as string in MongoDB
// 			if dob, err := time.Parse(time.RFC3339, dobStr); err == nil {
// 				if dob.Month() == currentMonth {
// 					upcoming = append(upcoming, m)
// 				}
// 			}
// 		}
// 	}

// 	c.JSON(http.StatusOK, upcoming)
// }

package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BirthdayHandler struct {
	Collection *mongo.Collection
}

func (h *BirthdayHandler) GetUpcomingBirthdays(c *gin.Context) {
	now := time.Now()
	currentMonth := int(now.Month()) // 1-12

	// MongoDB aggregation: project month of dateOfBirth and filter by current month
	pipeline := mongo.Pipeline{
		{{Key: "$addFields", Value: bson.M{
			"birthMonth": bson.M{"$month": "$dateOfBirth"},
		}}},
		{{Key: "$match", Value: bson.M{
			"birthMonth": currentMonth,
		}}},
	}

	cursor, err := h.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch birthdays"})
		return
	}

	var members []map[string]interface{}
	if err := cursor.All(context.TODO(), &members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse birthdays"})
		return
	}

	// Debug log
	log.Printf("Found %d birthdays in month %d\n", len(members), currentMonth)

	c.JSON(http.StatusOK, members)
}
