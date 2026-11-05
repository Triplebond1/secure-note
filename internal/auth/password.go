package auth

func HashPassword(password string) string {
	return "hashed-" + password
}

func validatePassword(password string) bool {
	return len(password) >= 8
}
