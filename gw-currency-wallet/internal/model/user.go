package model

// LogUser represents the login data for a user (typically used for authentication).
type LogUser struct {
	// Username is the user's login name.
	Username string `json:"username" binding:"required"`

	// Password is the user's password for authentication.
	Password string `json:"password" binding:"required"`
}

// User represents a user in the system, including their personal information and credentials.
type User struct {
	// Id is the unique identifier of the user in the system.
	Id int

	// Username is the user's chosen name for logging in.
	Username string `json:"username" binding:"required"`

	// Password is the user's encrypted password for secure authentication.
	Password string `json:"password" binding:"required"`

	// Email is the user's email address.
	Email string `json:"email" binding:"required"`
}
