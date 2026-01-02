package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"

	"fellowship-backend/internal/models"
)

type UserHandler struct {
	Collection *mongo.Collection
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		FullName   string `json:"fullName"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Role       string `json:"role"`
		SeniorCell string `json:"seniorCell"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	user := models.User{
		FullName:   input.FullName,
		Email:      input.Email,
		Password:   string(hashedPassword),
		Role:       input.Role,
		SeniorCell: input.SeniorCell,
		CreatedAt:  time.Now(),
	}

	_, err := h.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
