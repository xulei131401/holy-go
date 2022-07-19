package example

import (
	"context"
	"github.com/xulei131401/gox/responsex"

	"github.com/xulei131401/holy-go/service/api/admin/internal/svc"
	"github.com/xulei131401/holy-go/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DebugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDebugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DebugLogic {
	return &DebugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DebugLogic) Debug(req *types.DebugRequest) (resp *responsex.ApiResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
