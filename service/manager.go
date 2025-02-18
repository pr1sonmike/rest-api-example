package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pr1sonmike/rest-api-example/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	User        UserService
	FileMeta    FileMetaService
	FileContent FileContentService
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		User:        NewUserWebService(ctx, store),
		FileMeta:    NewFileMetaSvc(ctx, store),
		FileContent: NewFileContentSvc(ctx, store),
	}, nil
}
