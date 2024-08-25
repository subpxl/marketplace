package handlers

import (
	"marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AppConfig *config.AppConfig
}

func NewAuthHandler(appConfig *config.AppConfig) *AuthHandler {
	return &AuthHandler{AppConfig: appConfig}
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type SignupForm struct {
	Username        string `form:"username" binding:"required"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm-password" binding:"required,eqfield=Password"`
}

var users = map[string]string{
	"user1": "$2a$14$Uhz9JYlCoKxknD1P6t2W.eNTn.vYeFbxJD8LqP3RAC/U9AEoGip8O", // password: "password123"
}

func (h *AuthHandler) HandleLogin(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		c.HTML(200, "login.html", nil)

	} else {
		var form LoginForm

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, ok := users[form.Username]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(form.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		c.Redirect(http.StatusFound, "/dashboard")

	}

}
func (h *AuthHandler) HandleSignup(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(200, "signup.html", nil)

	} else {

		var form SignupForm

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, exists := users[form.Username]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Add the user to the map
		users[form.Username] = string(hashedPassword)

		c.Redirect(302, "/dashboard")
	}

}

func (h *AuthHandler) HandleLogout(c *gin.Context) {
	if c.Request.Method == http.MethodGet {

		c.JSON(200, "logout hit")
	}
}
