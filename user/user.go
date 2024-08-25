package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Email    string `json:"email"`
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
