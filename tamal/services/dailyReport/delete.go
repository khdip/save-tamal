package dailyreport

import (
	"context"
	"database/sql"

	dregrpc "save-tamal/proto/dailyReport"
	"save-tamal/tamal/storage"
)

func (s *Svc) DeleteDailyReport(ctx context.Context, req *dregrpc.DeleteDailyReportRequest) (*dregrpc.DeleteDailyReportResponse, error) {
	if err := s.drst.DeleteDailyReport(ctx, storage.DailyReport{
		ReportID: req.Dre.ReportID,
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{String: req.Dre.DeletedBy, Valid: true},
		},
	}); err != nil {
		return nil, err
	}

	return &dregrpc.DeleteDailyReportResponse{}, nil
}
