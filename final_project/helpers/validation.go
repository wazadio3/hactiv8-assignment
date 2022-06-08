package helpers

import (
	"final_project/models"

	"github.com/go-playground/validator/v10"
)

var validation *validator.Validate

func ValidatePayloadUser(data models.User) (map[string]interface{}, error) {
	invalidData := make(map[string]interface{})
	validation = validator.New()
	err := validation.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			invalidData[err.Field()] = err.Error()
		}
		return invalidData, err
	}
	return invalidData, nil
}

func CheckLoginPayload(data models.User) (map[string]interface{}, string) {
	invalidData := make(map[string]interface{})
	if data.Email == "" && data.Password == "" {
		invalidData["email"] = "required"
		invalidData["password"] = "required"
		return invalidData, "require email and password"
	} else if data.Email == "" {
		invalidData["email"] = "required"
		return invalidData, "require email"
	} else if data.Password == "" {
		invalidData["password"] = "required"
		return invalidData, "require password"
	}

	return invalidData, ""
}

func CheckUpdatePayload(data models.User) (map[string]interface{}, string) {
	invalidData := make(map[string]interface{})
	if data.Email == "" && data.UserName == "" {
		invalidData["email"] = "required"
		invalidData["username"] = "required"
		return invalidData, "require email and username"
	} else if data.Email == "" {
		invalidData["email"] = "required"
		return invalidData, "require email"
	} else if data.UserName == "" {
		invalidData["username"] = "required"
		return invalidData, "require username"
	}

	return invalidData, ""
}

func ValidatePayloadPhoto(data models.Photo) (map[string]interface{}, error) {
	invalidData := make(map[string]interface{})
	validation = validator.New()
	err := validation.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			invalidData[err.Field()] = err.Error()
		}
		return invalidData, err
	}
	return invalidData, nil
}

func ValidatePayloadComment(data models.Comment) (map[string]interface{}, error) {
	invalidData := make(map[string]interface{})
	validation = validator.New()
	err := validation.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			invalidData[err.Field()] = err.Error()
		}
		return invalidData, err
	}
	return invalidData, nil
}

func ValidatePayloadSosMed(data models.SosMed) (map[string]interface{}, error) {
	invalidData := make(map[string]interface{})
	validation = validator.New()
	err := validation.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			invalidData[err.Field()] = err.Error()
		}
		return invalidData, err
	}
	return invalidData, nil
}
