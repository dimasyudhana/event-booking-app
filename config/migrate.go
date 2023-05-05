package config

import (
	events "github.com/dimasyudhana/event-booking-app/app/features/events/repository"
	reviews "github.com/dimasyudhana/event-booking-app/app/features/reviews/repository"
	tickets "github.com/dimasyudhana/event-booking-app/app/features/tickets/repository"
	transactions "github.com/dimasyudhana/event-booking-app/app/features/transactions/repository"
	users "github.com/dimasyudhana/event-booking-app/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&users.User{},
		&events.Event{},
		&transactions.Transaction{},
		&transactions.Transaction_Tickets{},
		&transactions.Payment{},
		&tickets.Ticket{},
		&reviews.Review{},
	)
	return err
}
