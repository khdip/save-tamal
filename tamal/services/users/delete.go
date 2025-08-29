package users

import (
	"context"
	"database/sql"

	usergrpc "save-tamal/proto/users"
	"save-tamal/tamal/storage"
)

func (s *Svc) DeleteUser(ctx context.Context, req *usergrpc.DeleteUserRequest) (*usergrpc.DeleteUserResponse, error) {
	if err := s.ust.DeleteUser(ctx, storage.User{
		UserID: req.GetUser().UserID,
		CRUDTimeDate: storage.CRUDTimeDate{
			DeletedBy: sql.NullString{String: req.GetUser().DeletedBy, Valid: true},
		},
	}); err != nil {
		return nil, err
	}

	return &usergrpc.DeleteUserResponse{}, nil
}
