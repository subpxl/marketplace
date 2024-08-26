package user

import (
	"marketplace/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/render"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	AppConfig *config.AppConfig
}

func NewUserHandler(appConfig *config.AppConfig) *UserHandler {
	return &UserHandler{AppConfig: appConfig}
}

var users = map[string]string{
	"user1": "$2a$14$Uhz9JYlCoKxknD1P6t2W.eNTn.vYeFbxJD8LqP3RAC/U9AEoGip8O", // password: "password123"
}

func (h *UserHandler) HandleLogin(c *gin.Context) {

	if c.Request.Method == http.MethodGet {
		// c.HTML(200, "login.html", nil)

		if err := h.AppConfig.Render.HTML(c.Writer, http.StatusOK, "user/login", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
			Layout: "", // Disable layout
		}); err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
		}
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
func (h *UserHandler) HandleSignup(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		// c.HTML(200, "signup.html", nil)

		if err := h.AppConfig.Render.HTML(c.Writer, http.StatusOK, "user/signup", gin.H{"Title": "Dashboard"}, render.HTMLOptions{
			Layout: "", // Disable layout
		}); err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
		}
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

func (h *UserHandler) HandleLogout(c *gin.Context) {
	if c.Request.Method == http.MethodGet {

		c.JSON(200, "logout hit")
	}
}
