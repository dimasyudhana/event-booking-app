package main

import (
	"fmt"

	eventHandler "github.com/dimasyudhana/event-booking-app/app/features/events/handler"
	eventRepo "github.com/dimasyudhana/event-booking-app/app/features/events/repository"
	eventLogic "github.com/dimasyudhana/event-booking-app/app/features/events/services"
	reviewHandler "github.com/dimasyudhana/event-booking-app/app/features/reviews/handler"
	reviewRepo "github.com/dimasyudhana/event-booking-app/app/features/reviews/repository"
	reviewLogic "github.com/dimasyudhana/event-booking-app/app/features/reviews/services"
	ticketHandler "github.com/dimasyudhana/event-booking-app/app/features/tickets/handler"
	ticketRepo "github.com/dimasyudhana/event-booking-app/app/features/tickets/repository"
	ticketLogic "github.com/dimasyudhana/event-booking-app/app/features/tickets/services"
	transactionHandler "github.com/dimasyudhana/event-booking-app/app/features/transactions/handler"
	transactionRepo "github.com/dimasyudhana/event-booking-app/app/features/transactions/repository"
	transactionLogic "github.com/dimasyudhana/event-booking-app/app/features/transactions/services"
	userHandler "github.com/dimasyudhana/event-booking-app/app/features/users/handler"
	userRepo "github.com/dimasyudhana/event-booking-app/app/features/users/repository"
	userLogic "github.com/dimasyudhana/event-booking-app/app/features/users/service"
	"github.com/dimasyudhana/event-booking-app/app/routes"
	"github.com/dimasyudhana/event-booking-app/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg)
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)
	eventModel := eventRepo.New(db)
	eventService := eventLogic.New(eventModel)
	eventController := eventHandler.New(eventService)
	ticketModel := ticketRepo.New(db)
	ticketService := ticketLogic.New(ticketModel)
	ticketController := ticketHandler.New(ticketService)
	transactionModel := transactionRepo.New(db)
	transactionService := transactionLogic.New(transactionModel)
	transactionController := transactionHandler.New(transactionService)
	reviewModel := reviewRepo.New(db)
	reviewService := reviewLogic.New(reviewModel)
	reviewController := reviewHandler.New(reviewService)

	routes.Route(e, userController, eventController, ticketController, transactionController, reviewController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
