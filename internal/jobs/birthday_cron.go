package jobs

import (
	"os"
	"strings"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"

	"fellowship-backend/internal/services"
)

func StartBirthdayCron(
	memberCol *mongo.Collection,
	notificationCol *mongo.Collection,
) {
	adminEmails := strings.Split(os.Getenv("ADMIN_EMAILS"), ",")

	c := cron.New()
	c.AddFunc("0 6 * * *", func() {
		services.ProcessBirthdayNotifications(
			memberCol,
			notificationCol,
			adminEmails,
		)
	})
	c.Start()
}

