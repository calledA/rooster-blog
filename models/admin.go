package models

type Admin struct {
	ID       int
	Username    string
	Password string
}

func CheckAdmin(username, password string) bool {
	var admin Admin
	//其中&admin代表上面定义的结构体对象用于接收数据，使用时通过上面变量直接操作
	DB.Select("id").Where(Admin{Username:username,Password:password}).First(&admin)
	if admin.ID > 0 {
		return true
	}
	return false
}
