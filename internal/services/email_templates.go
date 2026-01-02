package services

import "fmt"

func BirthdayEmailTemplate(name, cell, notifType string, days int) string {
	var title string

	switch notifType {
	case "BIRTHDAY_7_DAYS":
		title = "🎉 Birthday in 7 Days"
	case "BIRTHDAY_1_DAY":
		title = "🎉 Birthday Tomorrow"
	case "BIRTHDAY_TODAY":
		title = "🎉 Birthday Today!"
	}

	return fmt.Sprintf(`
		<h2>%s</h2>
		<p><strong>%s</strong>'s birthday is in <strong>%d day(s)</strong>.</p>
		<p>Senior Cell: <strong>%s</strong></p>
		<p>Please prepare to celebrate 🎂</p>
	`, title, name, days, cell)
}
