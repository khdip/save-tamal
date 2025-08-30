package collection

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	collgrpc "save-tamal/proto/collection"
	"save-tamal/tamal/storage"
)

func (s *Svc) CreateCollection(ctx context.Context, req *collgrpc.CreateCollectionRequest) (*collgrpc.CreateCollectionResponse, error) {
	res, err := s.cst.CreateCollection(ctx, storage.Collection{
		CollectionID:  req.Coll.CollectionID,
		AccountType:   req.Coll.AccountType,
		AccountNumber: req.Coll.AccountNumber,
		Sender:        req.Coll.Sender,
		Date:          req.Coll.Date.AsTime(),
		Amount:        req.Coll.Amount,
		CRUDTimeDate: storage.CRUDTimeDate{
			CreatedBy: req.Coll.CreatedBy,
			UpdatedBy: req.Coll.UpdatedBy,
		},
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &collgrpc.CreateCollectionResponse{
		CollectionID: res,
	}, nil
}
