package entity

type SignUpParam struct {
	UserName string `json:"username" binding:"required"`
}
