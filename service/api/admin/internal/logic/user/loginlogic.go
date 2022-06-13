package user

import (
	"context"
	"errors"
	"holy-go/service/api/admin/internal/model"
	"holy-go/service/api/admin/internal/svc"
	"holy-go/service/api/admin/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Pwd != req.Password {
		return nil, errors.New("用户密码不正确")
	}

	return &types.LoginResponse{
		Id:     userInfo.Id,
		Name:   userInfo.Name,
		Gender: userInfo.Gender,
	}, nil
}
