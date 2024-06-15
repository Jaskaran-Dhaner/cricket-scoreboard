package validators

import (
	"accrualEngine/logger"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func ValidatePhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	logger.Log.With(zap.String("phone", phone)).Info(
		"validating phone number",
	)
	if len(phone) != 10 {
		logger.Log.With(
			zap.String("phone", phone)).Debug(
			"phone number is invalid",
		)
		return false
	}
	return true
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&()*+,-./:;<=>?@[\]^_{|}~]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	if !emailRegex.MatchString(email) {
		logger.Log.With(
			zap.String("email", email),
		).Debug("email is invalid")
		return false
	}
	return true
}

func ValidateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if len(name) < 3 || strings.TrimSpace(name) == "" {
		logger.Log.With(
			zap.String("name", name),
		).Debug("name is invalid")
		return false
	}
	return true
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 || strings.TrimSpace(password) == "" {
		return utf8.RuneCountInString(password) >= 8 && isASCII(password)
	}
	return true
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func ValidateLogin(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if len(username) < 3 || strings.TrimSpace(username) == "" {
		logger.Log.With(
			zap.String("username", username),
		).Debug("username is invalid")
		return false
	}
	return true
}

func ValidateAmount(fl validator.FieldLevel) bool {
	amount := fl.Field().String()
	if amount == "" || strings.TrimSpace(amount) == "" {
		return false
	}
	return true
}
