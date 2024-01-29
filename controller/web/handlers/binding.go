package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	validatorlib "github.com/go-playground/validator/v10"
)

func ShouldBind[T any](gc *gin.Context, obj T) error {
	vErr := new(validatorlib.ValidationErrors)
	if err := gc.ShouldBindHeader(obj); err != nil && !errors.As(err, vErr) {
		return err
	}
	if err := gc.ShouldBindUri(obj); err != nil && !errors.As(err, vErr) {
		return err
	}
	if err := gc.ShouldBind(obj); err != nil && !errors.As(err, vErr) {
		return err
	}
	if validator != nil {
		return validator.Struct(obj)
	}

	return nil
}
