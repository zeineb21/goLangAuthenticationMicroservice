package models

type User struct {
	UserId   string `form:"userid"`
	TenantId int 	`form:"tenantid"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
