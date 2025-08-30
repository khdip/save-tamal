package collection

import (
	"context"

	collgrpc "save-tamal/proto/collection"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Svc) GetUser(ctx context.Context, req *collgrpc.GetCollectionRequest) (*collgrpc.GetCollectionResponse, error) {
	r, err := s.cst.GetCollection(ctx, storage.Collection{
		CollectionID: req.Coll.CollectionID,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "collection doesn't exist")
	}
	return &collgrpc.GetCollectionResponse{
		Coll: &collgrpc.Collection{
			CollectionID:  r.CollectionID,
			AccountType:   r.AccountType,
			AccountNumber: r.AccountNumber,
			Sender:        r.Sender,
			Date:          timestamppb.New(r.Date),
			Amount:        r.Amount,
			CreatedAt:     timestamppb.New(r.CreatedAt),
			CreatedBy:     r.CreatedBy,
			UpdatedAt:     timestamppb.New(r.UpdatedAt),
			UpdatedBy:     r.UpdatedBy,
		},
	}, nil
}
