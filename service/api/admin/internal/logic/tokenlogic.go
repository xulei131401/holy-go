package logic

import (
	"context"
	"holy-go/service/api/admin/internal/svc"
	"holy-go/service/api/admin/internal/types"
	"holy-go/service/tool/jwt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TokenLogic {
	return &TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req *types.JwtTokenRequest) (resp *types.JwtTokenResponse, err error) {
	// todo: add your logic here and delete this line
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire

	now := time.Now().Unix()
	accessToken, err := jwt.GetToken(now, l.svcCtx.Config.JwtAuth.AccessSecret, nil, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
