package reviews

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	UserID   uint
	Username string
	Image    string
	EventID  uint
	Review   string
}

type Repository interface {
	WriteReview(Core) (Core, error)
	UpdateReview(Core) (Core, error)
	DeleteReview(id uint) error
}

type Service interface {
	WriteReview(Core) (Core, error)
	UpdateReview(Core) (Core, error)
	DeleteReview(id uint) error
}

type Handler interface {
	WriteReview() echo.HandlerFunc
	UpdateReview() echo.HandlerFunc
	DeleteReview() echo.HandlerFunc
}
