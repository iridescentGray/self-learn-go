package logic

import (
	"context"

	"fisrt/internal/svc"
	"fisrt/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FisrtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFisrtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FisrtLogic {
	return &FisrtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FisrtLogic) Fisrt(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "My Test Msg"
	return
}
