package models

type Admin struct {
	ID       int
	Email    string
	Password string
}

func CheckAdmin(email, password string) bool {
	var admin Admin
	DB.Select("id").Where(Admin{Email:email,Password:password}).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}
