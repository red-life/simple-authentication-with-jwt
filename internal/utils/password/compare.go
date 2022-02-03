package password

func ComparePasswords(password string, hashedPassword string) bool{
	hashPassword := HashPassword(password)
	if hashedPassword != hashPassword{
		return false
	}
	return true
}