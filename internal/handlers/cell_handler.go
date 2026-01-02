package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CellHandler struct {
	Collection *mongo.Collection
}

func (h *CellHandler) GetSeniorCellsSummary(c *gin.Context) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$fellowship.seniorCell"},
			{Key: "memberCount", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "seniorCell", Value: "$_id"},
			{Key: "memberCount", Value: 1},
		}}},
	}

	cursor, err := h.Collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

func (h *CellHandler) GetMembersBySeniorCell(c *gin.Context) {
	seniorCell := c.Param("seniorCell")

	filter := bson.M{
		"fellowship.seniorCell": seniorCell,
	}

	cursor, err := h.Collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var members []bson.M
	if err := cursor.All(context.TODO(), &members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

