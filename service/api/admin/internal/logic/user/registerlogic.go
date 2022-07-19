package user

import (
	"context"
	"errors"

	"github.com/xulei131401/gox/responsex"
	"github.com/xulei131401/holy-go/service/api/admin/internal/model"
	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *responsex.ApiResponse, err error) {
	// todo: add your logic here and delete this line

	stu, _ := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if stu != nil {
		return nil, errors.New("账号已存在，请重新输入")
	}

	user := &model.User{
		Username: req.Username,
		Name:     req.Name,
		Pwd:      req.Password,
	}

	res, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	affected, _ := res.RowsAffected()
	if affected != 1 {
		return nil, errors.New("插入失败")
	}

	return responsex.NewSuccessApiResponse(), nil
}
