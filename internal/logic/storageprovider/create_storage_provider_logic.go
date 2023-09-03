package storageprovider

import (
	"context"
	"github.com/suyuan32/simple-admin-file/internal/utils/cloudsdk"

	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"
	"github.com/suyuan32/simple-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStorageProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStorageProviderLogic {
	return &CreateStorageProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStorageProviderLogic) CreateStorageProvider(req *types.StorageProviderInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.StorageProvider.Create().
		SetNotNilState(req.State).
		SetNotNilName(req.Name).
		SetNotNilBucket(req.Bucket).
		SetNotNilProviderName(req.ProviderName).
		SetNotNilSecretID(req.SecretId).
		SetNotNilSecretKey(req.SecretKey).
		SetNotNilRegion(req.Region).
		SetNotNilIsDefault(req.IsDefault).
		SetNotNilFolder(req.Folder).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudUploader = cloudsdk.NewUploaderGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
