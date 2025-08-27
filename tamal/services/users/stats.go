package users

import (
	"context"

	usergrpc "save-tamal/proto/users"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) UserStats(ctx context.Context, req *usergrpc.UserStatsRequest) (*usergrpc.UserStatsResponse, error) {
	r, err := h.ust.UserStats(ctx, storage.Filter{
		SearchTerm: req.Filter.SearchTerm,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "user doesn't exist")
	}
	return &usergrpc.UserStatsResponse{
		Stats: &usergrpc.Stats{
			Count:       r.Count,
			TotalAmount: r.TotalAmount,
		},
	}, nil
}
