package collection

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) CollectionStats(ctx context.Context, filter storage.Filter) (storage.Stats, error) {
	collStats, err := s.st.CollectionStats(ctx, filter)
	if err != nil {
		return storage.Stats{}, status.Error(codes.Internal, "processing failed")
	}

	return collStats, nil
}
