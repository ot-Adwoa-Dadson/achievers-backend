package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"fellowship-backend/internal/models"
)

func ProcessBirthdayNotifications(
	memberCol *mongo.Collection,
	notificationCol *mongo.Collection,
	adminEmails []string,
) {
	cursor, err := memberCol.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return
	}

	var members []models.Member
	if err := cursor.All(context.TODO(), &members); err != nil {
		log.Println(err)
		return
	}

	for _, member := range members {
		days := daysUntilBirthday(member.DateOfBirth)

		var notifType string
		switch days {
		case 7:
			notifType = "BIRTHDAY_7_DAYS"
		case 1:
			notifType = "BIRTHDAY_1_DAY"
		case 0:
			notifType = "BIRTHDAY_TODAY"
		default:
			continue
		}

		year := time.Now().Year()

		// Check if notification already sent
		count, _ := notificationCol.CountDocuments(
			context.TODO(),
			bson.M{
				"memberId": member.ID,
				"type":     notifType,
				"year":     year,
			},
		)

		if count > 0 {
			continue
		}

		emailBody := BirthdayEmailTemplate(
			member.FullName,
			member.Fellowship.SeniorCell,
			notifType,
			days,
		)

		err := SendEmail(
			adminEmails,
			"🎉 Birthday Reminder",
			emailBody,
		)

		if err != nil {
			log.Println("Email failed:", err)
			continue
		}

		// Log notification
		notificationCol.InsertOne(context.TODO(), bson.M{
			"memberId": member.ID,
			"type":     notifType,
			"year":     year,
			"sentAt":   time.Now(),
		})

		log.Printf("Sent %s notification for %s\n", notifType, member.FullName)
	}
}


func daysUntilBirthday(dob time.Time) int {
	now := time.Now()

	thisYearBirthday := time.Date(
		now.Year(),
		dob.Month(),
		dob.Day(),
		0, 0, 0, 0,
		time.Local,
	)

	if thisYearBirthday.Before(now) {
		thisYearBirthday = thisYearBirthday.AddDate(1, 0, 0)
	}

	return int(thisYearBirthday.Sub(now).Hours() / 24)
}
