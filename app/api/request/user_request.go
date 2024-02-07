package request

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

type PasswordChangeRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
