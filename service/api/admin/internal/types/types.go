// Code generated by goctl. DO NOT EDIT.
package types

type EmptyReply struct {
}

type DebugRequest struct {
}

type TestRequest struct {
}

type JwtTokenRequest struct {
}

type JwtTokenResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type LoginRequest struct {
	Username string `form:"username,optional" validate:"required"`
	Password string `form:"password,optional" validate:"required"`
}

type LoginResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RegisterRequest struct {
	Username string `form:"username,optional" validate:"required"`
	Name     string `form:"name,optional" validate:"required"`
	Password string `form:"password,optional" validate:"required"`
}

type UserInfoRequest struct {
	Id int64 `form:"id,optional" validate:"required"`
}

type UserInfoResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
