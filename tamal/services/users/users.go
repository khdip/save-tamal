package users

import (
	"context"

	"google.golang.org/grpc"

	usergrpc "save-tamal/proto/users"
	"save-tamal/tamal/storage"
)

type UserStore interface {
	CreateUser(ctx context.Context, ust storage.User) (string, error)
	GetUser(ctx context.Context, ust storage.User) (*storage.User, error)
	UpdateUser(ctx context.Context, ust storage.User) (*storage.User, error)
	DeleteUser(ctx context.Context, ust storage.User) error
	ListUser(ctx context.Context, flt storage.Filter) ([]storage.User, error)
	UserStats(ctx context.Context, flt storage.Filter) (storage.Stats, error)
}

type Handler struct {
	usergrpc.UnimplementedUserServiceServer
	ust UserStore
}

func New(cs UserStore) *Handler {
	return &Handler{
		ust: cs,
	}
}

// RegisterService with grpc server.
func (h *Handler) RegisterSvc(srv *grpc.Server) error {
	usergrpc.RegisterUserServiceServer(srv, h)
	return nil
}
