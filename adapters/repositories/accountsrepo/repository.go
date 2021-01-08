package accountsrepo

import (
	"example.com/test_task/core/ports"
	"example.com/test_task/pkg"
)

type repository struct {
	storageClient pkg.StorageClient
}

func New(storageClient pkg.StorageClient) ports.AccountsRepository {
	return &repository{storageClient: storageClient}
}