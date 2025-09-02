package dailyreport

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) CreateDailyReport(ctx context.Context, dre storage.DailyReport) (int32, error) {
	rid, err := s.st.CreateDailyReport(ctx, dre)
	if err != nil {
		return 0, status.Error(codes.Internal, "processing failed")
	}

	return rid, nil
}
