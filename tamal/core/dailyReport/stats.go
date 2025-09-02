package dailyreport

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) DailyReportStats(ctx context.Context, filter storage.Filter) (storage.Stats, error) {
	drStats, err := s.st.DailyReportStats(ctx, filter)
	if err != nil {
		return storage.Stats{}, status.Error(codes.Internal, "processing failed")
	}

	return drStats, nil
}
