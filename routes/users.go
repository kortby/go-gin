package routes

import (
	"database/sql"
	"example/gingonic/db"
	"example/gingonic/models"
	"example/gingonic/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": "Could not parse data", "details": err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		context.JSON(500, gin.H{"message": "Could not create user. Try again later.", "details": err.Error()})
		return
	}
	context.JSON(201, gin.H{"message": "User Created", "user": user})
}

func loginHandler(c *gin.Context) {
	var credentials LoginCredentials
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	// Use the db.DB variable from the db package that's initialized on app startup
	row := db.DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", credentials.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !verifyPassword(user.Password, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle successful login, such as generating and sending a token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

type LoginCredentials struct {
	Email    string `json:"email"` // Adjusted to match the JSON payload
	Password string `json:"password"`
}

func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
