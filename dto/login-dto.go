package dto

//LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

// URL	jdbc:mysql://localhost:3306/facilityManage charset=%s&parseTime=%s&loc=%s
// migrate -path db/migration -database "mysql://facilityManage:123456@tcp(localhost:3306)/facilityManage?charset=utf8mb4&parseTime=true&loc=Local" -verbose down
