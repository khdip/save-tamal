package collection

import (
	"context"
	"database/sql"

	collgrpc "save-tamal/proto/collection"
	"save-tamal/tamal/storage"
)

func (s *Svc) DeleteCollection(ctx context.Context, req *collgrpc.DeleteCollectionRequest) (*collgrpc.DeleteCollectionResponse, error) {
	if err := s.cst.DeleteCollection(ctx, storage.Collection{
		CollectionID: req.Coll.CollectionID,
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{String: req.Coll.DeletedBy, Valid: true},
		},
	}); err != nil {
		return nil, err
	}

	return &collgrpc.DeleteCollectionResponse{}, nil
}
