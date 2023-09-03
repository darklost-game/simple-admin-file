package storageprovider

import (
	"context"

	"github.com/suyuan32/simple-admin-file/ent/predicate"
	"github.com/suyuan32/simple-admin-file/ent/storageprovider"
	"github.com/suyuan32/simple-admin-file/internal/svc"
	"github.com/suyuan32/simple-admin-file/internal/types"
	"github.com/suyuan32/simple-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorageProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStorageProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorageProviderListLogic {
	return &GetStorageProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStorageProviderListLogic) GetStorageProviderList(req *types.StorageProviderListReq) (*types.StorageProviderListResp, error) {
	var predicates []predicate.StorageProvider
	if req.Name != nil {
		predicates = append(predicates, storageprovider.NameContains(*req.Name))
	}
	if req.Bucket != nil {
		predicates = append(predicates, storageprovider.BucketContains(*req.Bucket))
	}
	if req.ProviderName != nil {
		predicates = append(predicates, storageprovider.ProviderNameContains(*req.ProviderName))
	}
	data, err := l.svcCtx.DB.StorageProvider.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.StorageProviderListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.StorageProviderInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.Unix()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.Unix()),
				},
				State:        &v.State,
				Name:         &v.Name,
				Bucket:       &v.Bucket,
				ProviderName: &v.ProviderName,
				SecretId:     &v.SecretID,
				SecretKey:    &v.SecretKey,
				Region:       &v.Region,
				IsDefault:    &v.IsDefault,
				Folder:       &v.Folder,
			})
	}

	return resp, nil
}
