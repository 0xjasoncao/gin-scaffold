package dto

// LoginRequest 登录请求参数
type LoginRequest struct {
	Password   string `json:"password,omitempty" binding:"required"`
	VerifyCode string `json:"verify-code,omitempty" binding:"required"`
	Email      string `json:"email,omitempty" binding:"required,email"`
}
