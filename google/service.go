package transfer

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	transfer "google.golang.org/api/admin/datatransfer/v1"
	directory "google.golang.org/api/admin/directory/v1"
)

type Service struct {
	TransferClient  *transfer.Service
	DirectoryClient *directory.Service
}

func NewService(credentials, delegate string) (*Service, error) {
	tc, err := CreateTransferClient(credentials, delegate)
	if err != nil {
		return nil, err
	}

	dc, err := CreateDirectoryClient(credentials, delegate)
	if err != nil {
		return nil, err
	}

	s := &Service{
		TransferClient:  tc,
		DirectoryClient: dc,
	}

	return s, nil
}

func CreateTransferClient(credentials, delegate string) (*transfer.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(credentials)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(b, transfer.AdminDatatransferScope)
	if err != nil {
		return nil, err
	}
	config.Subject = delegate
	ts := config.TokenSource(ctx)

	srv, err := transfer.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func CreateDirectoryClient(credentials, delegate string) (*directory.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(credentials)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(b, directory.AdminDirectoryUserScope)
	if err != nil {
		return nil, err
	}
	config.Subject = delegate
	ts := config.TokenSource(ctx)

	srv, err := directory.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return srv, nil
}
