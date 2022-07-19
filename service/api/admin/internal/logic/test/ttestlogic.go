package test

import (
	"context"

	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xulei131401/gox/responsex"
)

type TTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TTestLogic {
	return &TTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TTestLogic) TTest(req *types.TestRequest) (resp *responsex.ApiResponse, err error) {
	// todo: add your logic here and delete this line

	return nil, nil
}
