package collection

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) UpdateCollection(ctx context.Context, coll storage.Collection) (*storage.Collection, error) {
	c, err := s.st.UpdateCollection(ctx, coll)
	if err != nil {
		return nil, status.Error(codes.Internal, "processing failed")
	}

	return &storage.Collection{
		CollectionID:  c.CollectionID,
		AccountType:   c.AccountType,
		AccountNumber: c.AccountNumber,
		Sender:        c.Sender,
		Date:          c.Date,
		Amount:        c.Amount,
		CRUDTimeDate: storage.CRUDTimeDate{
			UpdatedBy: c.UpdatedBy,
		},
	}, nil
}
