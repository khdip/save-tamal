package users

import (
	"context"

	usergrpc "save-tamal/proto/users"
	"save-tamal/tamal/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *Handler) GetUser(ctx context.Context, req *usergrpc.GetUserRequest) (*usergrpc.GetUserResponse, error) {
	r, err := h.ust.GetUser(ctx, storage.User{
		UserID: req.GetUser().UserID,
		Email:  req.GetUser().Email,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, "user doesn't exist")
	}
	return &usergrpc.GetUserResponse{
		User: &usergrpc.User{
			UserID:    r.UserID,
			Name:      r.Name,
			Batch:     r.Batch,
			Email:     r.Email,
			Password:  r.Password,
			CreatedAt: timestamppb.New(r.CreatedAt),
			CreatedBy: r.CreatedBy,
			UpdatedAt: timestamppb.New(r.UpdatedAt),
			UpdatedBy: r.UpdatedBy,
		},
	}, nil
}
