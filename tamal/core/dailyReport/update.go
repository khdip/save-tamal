package dailyreport

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) UpdateDailyReport(ctx context.Context, dre storage.DailyReport) (*storage.DailyReport, error) {
	dr, err := s.st.UpdateDailyReport(ctx, dre)
	if err != nil {
		return nil, status.Error(codes.Internal, "processing failed")
	}

	return dr, nil
}
