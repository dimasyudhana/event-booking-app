package handler

import (
	"net/http"
	"strconv"

	"github.com/dimasyudhana/event-booking-app/app/features/tickets"
	"github.com/dimasyudhana/event-booking-app/helper"
	"github.com/dimasyudhana/event-booking-app/middlewares"
	"github.com/labstack/echo/v4"
)

type TicketController struct {
	s tickets.Service
}

func New(h tickets.Service) tickets.Handler {
	return &TicketController{s: h}
}

func (tc *TicketController) GetTickets() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		tickets, err := tc.s.GetTickets(uint(id))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		var response []ResponseGetTickets
		for _, ticket := range tickets {
			response = append(response, ResponseGetTickets{
				EventID:        ticket.EventID,
				TicketID:       ticket.ID,
				TicketCategory: ticket.TicketCategory,
				TicketPrice:    ticket.TicketPrice,
				TicketQuantity: ticket.TicketQuantity,
			})
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    response,
		})
	}
}

func (tc *TicketController) UpdateTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputs []RequestUpdateTicket
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		event_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&inputs); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		var updatedTickets []tickets.Core
		for _, input := range inputs {
			updatedTickets = append(updatedTickets, tickets.Core{
				TicketCategory: input.TicketCategory,
				TicketPrice:    input.TicketPrice,
				TicketQuantity: input.TicketQuantity,
				EventID:        uint(event_id),
			})
		}

		err = tc.s.UpdateTicket(uint(event_id), updatedTickets)
		if err != nil {
			c.Logger().Error("Failed to update ticket: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Updated Tickets", nil))
	}
}

func (tc *TicketController) DeleteTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		err = tc.s.DeleteTicket(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Deleted a Tickets", nil))
	}
}
