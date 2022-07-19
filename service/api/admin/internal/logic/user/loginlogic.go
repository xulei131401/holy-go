package user

import (
	"context"
	"errors"
	"github.com/xulei131401/gox/jwt"
	"github.com/xulei131401/gox/responsex"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"github.com/xulei131401/holy-go/service/api/admin/internal/model"
	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *responsex.ApiResponse, err error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, err
	}

	if userInfo.Pwd != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	loginResponse := &types.LoginResponse{
		Id:   userInfo.Id,
		Name: userInfo.Name,
	}

	token, err := l.Token(userInfo)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m["user"] = loginResponse
	m["token"] = token
	return responsex.NewSuccessApiResponseWithData(m), nil
}

func (l *LoginLogic) Token(user *model.User) (resp *types.JwtTokenResponse, err error) {
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	now := time.Now().Unix()
	payloads := make(map[string]interface{})
	if user != nil {
		payloads["uid"] = user.Id
	}

	accessToken, err := jwt.GetToken(now, l.svcCtx.Config.Auth.AccessSecret, payloads, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
