package logic

import (
	"context"
	"github.com/zeromicro/x/errors"
	"wechat-ocr/api/internal/svc"
	"wechat-ocr/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PickupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPickupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PickupInfoLogic {
	return &PickupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PickupInfoLogic) PickupInfo(req *types.PickupInfoRequest) (resp *types.PickupInfoResponse, err error) {
	// todo: add your logic here and delete this line
	if req.Imagescode == "" {
		return nil, errors.New(1001, "请求参数不能为空")
	}
	//resultMap, err := logic.SPickUpInfo(req.Imagescode)
	//if err != nil {
	//	return nil, errors.New(1002, "获取信息失败")
	//}

	return
}
