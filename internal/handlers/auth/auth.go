package auth

import (
	"net/http"
	"scoreboard/config"
	"scoreboard/internal/models/auth"
	authUtils "scoreboard/internal/utils/auth"
	"scoreboard/internal/utils/database"
	"scoreboard/internal/validators"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupAuthHandlers(router *gin.RouterGroup) {
	// Public routes
	router.POST("/signup", SignupHandler)
	router.POST("/login", LoginHandler)
}

// SignupHandler godoc
// @Summary Sign up a new user
// @Description Creates a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.UserRegister true "Register User"
// @Router /public/auth/signup [post]
func SignupHandler(c *gin.Context) {

	// Bind and validate the JSON payload
	var request auth.UserRegister
	if err := validators.BindAndValidate(c, &request); err != nil {
		return
	}

	// Hash the password
	hashedPassword, err := authUtils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new user
	user := auth.User{
		Name:      request.Name,
		Email:     request.Email,
		Password:  hashedPassword,
		Role:      auth.UserRole,
		Gender:    request.Gender,
		IsActive:  true,
		IsDeleted: false,
		CreatedBy: "system",
		UpdatedBy: "system",
		Version:   config.AppConfig.Version,
	}

	exists, err := database.GetCollection(config.AppConfig.Collections.Users).CountDocuments(c, bson.M{"email": request.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		return
	}

	// Save the user to the database
	_, err = database.GetCollection(config.AppConfig.Collections.Users).InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})

}

// LoginHandler godoc
// @Summary Login a user
// @Description Logs in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body auth.UserLogin true "Login User"
// @Success 200 {object} auth.LoginResponse
// @Router /public/auth/login [post]
func LoginHandler(c *gin.Context) {
	var loginRequest auth.UserLogin
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user in the database
	var user auth.User
	err := database.GetCollection(config.AppConfig.Collections.Users).FindOne(c, bson.M{"email": loginRequest.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Check if the password is correct
	verified := authUtils.VerifyPassword(loginRequest.Password, user.Password)

	if !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate access and refresh tokens
	accessToken, refreshToken, err := authUtils.GenerateTokens(user.Name, user.Email, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the access and refresh tokens
	c.JSON(http.StatusOK, auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role:         user.Role,
		Email:        user.Email,
		Name:         user.Name,
	})
}
