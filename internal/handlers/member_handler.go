package handlers

import (
	"context"
	"net/http"

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
