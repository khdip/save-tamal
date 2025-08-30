package collection

import "save-tamal/tamal/storage/postgres"

type CoreSvc struct {
	st *postgres.Storage
}

func New(st *postgres.Storage) *CoreSvc {
	return &CoreSvc{
		st: st,
	}
}
