package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" `
	Phone    string `json:"phone" binding:"required" `
	RoleID   uint64 `json:"role_id" form:"id"`
	Password string `json:"password" form:"password" binding:"required"`
}
