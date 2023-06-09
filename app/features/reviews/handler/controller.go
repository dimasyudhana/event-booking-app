package handler

import (
	"net/http"
	"strconv"

	"github.com/dimasyudhana/event-booking-app/app/features/reviews"
	"github.com/dimasyudhana/event-booking-app/helper"
	"github.com/dimasyudhana/event-booking-app/middlewares"
	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	n reviews.Service
}

func New(o reviews.Service) reviews.Handler {
	return &ReviewController{n: o}
}

func (rc *ReviewController) WriteReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestWriteReview
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userid := claims.ID
		username := claims.Username
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		request := reviews.Core{
			UserID:   userid,
			Username: username,
			EventID:  uint(eventid),
			Review:   input.Review,
		}

		_, err = rc.n.WriteReview(request)
		if err != nil {
			c.Logger().Error("Failed to write a review:", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created a Review", nil))
	}
}

func (rc *ReviewController) UpdateReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateReview
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userid := claims.ID
		username := claims.Username
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input:", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		request := reviews.Core{
			UserID:   userid,
			Username: username,
			EventID:  uint(eventid),
			Review:   input.Review,
		}

		_, err = rc.n.UpdateReview(request)
		if err != nil {
			c.Logger().Error("Failed to update review: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))

		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Updated a Review", nil))
	}
}

func (rc *ReviewController) DeleteReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if claims.ID != uint(id) {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is not valid for this user.", nil))
		}

		err = rc.n.DeleteReview(uint(id))
		if err != nil {
			c.Logger().Error("Failed to delete review: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success Deleted a Review", nil))
	}
}
