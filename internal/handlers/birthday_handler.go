package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BirthdayHandler struct {
	Collection *mongo.Collection
}

// GetUpcomingBirthdays returns members with birthdays in the next 7 days
func (h *BirthdayHandler) GetUpcomingBirthdays(c *gin.Context) {
	now := time.Now()
	day7 := now.AddDate(0, 0, 7)

	// Filter by birthday (ignore year)
	filter := bson.M{
		"dateOfBirth": bson.M{
			"$gte": now,
			"$lte": day7,
		},
	}

	cursor, err := h.Collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch birthdays"})
		return
	}
	var members []map[string]interface{}
	if err := cursor.All(context.TODO(), &members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse birthdays"})
		return
	}

	c.JSON(http.StatusOK, members)
}
