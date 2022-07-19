package user

import (
	"context"
	"errors"
	"github.com/xulei131401/holy-go/service/api/admin/internal/model"

	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xulei131401/gox/responsex"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.UserInfoRequest) (resp *responsex.ApiResponse, err error) {
	logx.Infof("userId: %v", l.ctx.Value("uid"))
	userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户不存在")
	default:
		return nil, err
	}

	loginResponse := &types.LoginResponse{
		Id:   userInfo.Id,
		Name: userInfo.Name,
	}

	return responsex.NewSuccessApiResponseWithData(loginResponse), nil
}
