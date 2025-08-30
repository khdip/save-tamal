package collection

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) GetCollection(ctx context.Context, coll storage.Collection) (*storage.Collection, error) {
	c, err := s.st.GetCollection(ctx, coll)
	if err != nil {
		return nil, status.Error(codes.Internal, "processing failed")
	}
	return c, nil
}
