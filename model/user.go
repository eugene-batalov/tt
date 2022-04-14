package model

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	User struct {
		ID       int64   `gorm:"primaryKey" json:"id"`
		FullName string  `json:"fullName" gorm:"column:fullName" validate:"required"`
		Email    string  `json:"email" validate:"required,email"`
		Phone    *string `json:"phone"`
		Age      *int    `json:"age"`
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
