package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"fellowship-backend/internal/models"
)

type MemberHandler struct {
	Collection *mongo.Collection
}

func (h *MemberHandler) GetAllMembers(c *gin.Context) {
	var members []models.Member

	cursor, err := h.Collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cursor.All(context.TODO(), &members); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *MemberHandler) CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()

	_, err := h.Collection.InsertOne(context.Background(), member)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create member"})
		return
	}

	c.JSON(http.StatusCreated, member)
}