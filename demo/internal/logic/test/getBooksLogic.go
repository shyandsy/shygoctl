package test

import (
	"context"

	"github.com/shyandsy/shygoctl/demo/internal/svc"
	"github.com/shyandsy/shygoctl/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBooksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get books
func NewGetBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBooksLogic {
	return &GetBooksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBooksLogic) GetBooks(req *types.GetBookReq) (resp *types.GetBookResp, err error) {
	// todo: add your logic here and delete this line

	return
}
