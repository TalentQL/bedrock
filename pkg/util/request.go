package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Parse[T any](c *gin.Context) (T, error) {
	var payload T

	// Bind JSON
	if err := c.ShouldBindJSON(&payload); err != nil {
		return payload, err
	}

	// Validate
	if err := validate.Struct(payload); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok && len(errs) > 0 {
			// Return only the first validation error
			return payload, errs[0]
		}
		return payload, err
	}

	return payload, nil
}
