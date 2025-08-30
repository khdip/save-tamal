package collection

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) CreateCollection(ctx context.Context, coll storage.Collection) (string, error) {
	collid, err := s.st.CreateCollection(ctx, coll)
	if err != nil {
		return "", status.Error(codes.Internal, "processing failed")
	}

	return collid, nil
}
