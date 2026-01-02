package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"fellowship-backend/internal/handlers"
)

func RegisterRoutes(router *gin.Engine, memberCol, userCol *mongo.Collection) {
	memberHandler := handlers.MemberHandler{Collection: memberCol}
	userHandler := handlers.UserHandler{Collection: userCol}
	cellHandler := handlers.CellHandler{Collection: memberCol}
	birthdayHandler := handlers.BirthdayHandler{Collection: memberCol}

	api := router.Group("/api")
	{
		api.GET("/members", memberHandler.GetAllMembers)
		api.POST("/users", userHandler.CreateUser)

		api.GET("/cells", cellHandler.GetSeniorCellsSummary)
		api.GET("/cells/:seniorCell/members", cellHandler.GetMembersBySeniorCell)
		api.GET("/birthdays/upcoming", birthdayHandler.GetUpcomingBirthdays)
	}
}
