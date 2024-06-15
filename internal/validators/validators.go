package validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("phone", ValidatePhone)
	validate.RegisterValidation("amount", ValidateAmount)
	validate.RegisterValidation("password", ValidatePassword)
	validate.RegisterValidation("email", ValidateEmail)
	validate.RegisterValidation("name", ValidateName)
	validate.RegisterValidation("login", ValidateLogin)
}

// BindAndValidate binds the JSON payload to the struct and validates it.
func BindAndValidate(c *gin.Context, obj interface{}) error {
	// Bind the JSON payload to the struct
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// Validate the struct
	err := validate.Struct(obj)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}

		validationErrors := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, err := range validationErrors {
			errors[err.Field()] = err.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return err
	}

	return nil
}
