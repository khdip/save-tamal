package dailyreport

import (
	"context"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CoreSvc) DeleteDailyReport(ctx context.Context, dre storage.DailyReport) error {

	if err := s.st.DeleteDailyReport(ctx, dre); err != nil {
		return status.Error(codes.Internal, "processing failed")
	}
	return nil
}
