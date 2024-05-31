package utils

import (
	"encoding/base64"
	"net/http"

	"errors"

	"github.com/go-playground/validator/v10"
)

func SchemaValidator[T any](req *T) (int, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(req)

	if err != nil {
		err_validator := err.(validator.ValidationErrors)
		return http.StatusBadRequest, errors.New(err_validator.Error())
	}

	return http.StatusOK, nil
}

func Base64Validate(dat []byte) error {

	_, err := base64.StdEncoding.DecodeString(string(dat))
	if err != nil {
		return errors.New("FileData must be a valid base64 string")
	}

	return nil
}
