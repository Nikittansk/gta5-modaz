package user

type user struct {
	ID       int    `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
