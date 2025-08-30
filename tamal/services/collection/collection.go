package collection

import (
	"context"

	"google.golang.org/grpc"

	collgrpc "save-tamal/proto/collection"
	"save-tamal/tamal/storage"
)

type CollectionStore interface {
	CreateCollection(ctx context.Context, cst storage.Collection) (string, error)
	GetCollection(ctx context.Context, cst storage.Collection) (*storage.Collection, error)
	UpdateCollection(ctx context.Context, cst storage.Collection) (*storage.Collection, error)
	DeleteCollection(ctx context.Context, cst storage.Collection) error
	ListCollection(ctx context.Context, flt storage.Filter) ([]storage.Collection, error)
	CollectionStats(ctx context.Context, flt storage.Filter) (storage.Stats, error)
}

type Svc struct {
	collgrpc.UnimplementedCollectionServiceServer
	cst CollectionStore
}

func New(cs CollectionStore) *Svc {
	return &Svc{
		cst: cs,
	}
}

// RegisterService with grpc server.
func (s *Svc) RegisterSvc(srv *grpc.Server) error {
	collgrpc.RegisterCollectionServiceServer(srv, s)
	return nil
}
