package comments

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) ListComment(ctx context.Context, filter storage.Filter) ([]storage.Comment, error) {
	lst, err := s.st.ListComment(ctx, filter)
	if err != nil {
		return nil, status.Error(codes.Internal, "processing failed")
	}

	return lst, nil
}
