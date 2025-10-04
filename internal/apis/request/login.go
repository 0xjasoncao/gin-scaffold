package request

// LoginRequest 登录
type LoginRequest struct {
	Password   string `json:"password,omitempty" binding:"required,min=6,max=16"`
	VerifyCode string `json:"verify_code,omitempty" binding:"required"`
	Mobile     string `json:"mobile" binding:"required,len=11"`
}
