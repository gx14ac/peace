package entity

type SignUpParam struct {
	UserName string `json:"username" binding:"required"`
}

type UpdateParam struct {
	UserName string `json:"username" Binding:"required"`
}
