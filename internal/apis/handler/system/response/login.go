package response

import "gin-scaffold/pkg/token"

type LoginResponse struct {
	Id        string                  `json:"id,omitempty"`
	UserName  string                  `json:"user_name,omitempty"`
	Gender    int                     `json:"gender,omitempty"`
	NickName  string                  `json:"nick_name,omitempty"`
	TokenInfo *token.IssuingTokenInfo `json:"token_info"`
}
