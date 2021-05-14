// Copyright (c) 2021, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	sentryecho "github.com/getsentry/sentry-go/echo"
	echo "github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
)

func httpErrorHandler(err error, c echo.Context) {
	if err == mongo.ErrNoDocuments {
		if jsonErr := c.JSON(http.StatusNotFound, newErrorResponse("Not Found")); jsonErr != nil {
			c.Logger().Error(jsonErr)
		}
		return
	}

	if c.Response().Committed {
		return
	}

	c.Logger().Error(err)

	if _, ok := err.(validator.ValidationErrors); ok {
		jsonErr := c.JSON(http.StatusUnprocessableEntity, newErrorResponse(err.Error()))
		if jsonErr != nil {
			c.Logger().Error(jsonErr)
		}
		return
	}

	httpError, isHTTPError := errors.Cause(err).(*echo.HTTPError)
	if isHTTPError {
		jsonErr := c.JSON(httpError.Code, newErrorResponse(err.Error()))
		if jsonErr != nil {
			c.Logger().Error(jsonErr)
		}
		return
	}

	jsonErr := c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	if jsonErr != nil {
		c.Logger().Error(jsonErr)
	}

	if hub := sentryecho.GetHubFromContext(c); hub != nil {
		hub.CaptureException(err)
	}
}

type errorResponse struct {
	Errors map[string]string `json:"errors"`
}

func newErrorResponse(message string) *errorResponse {
	return &errorResponse{
		Errors: map[string]string{
			"_all": message,
		},
	}
}
