package initstorage

import (
	"context"
	"fmt"

	"github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/kristina71/avito_otus/hw12_13_14_15_calendar/internal/storage/sql"
)

func NewStorage(ctx context.Context, storageType string, dsn string) (storage.Storage, error) {
	var db storage.Storage
	switch storageType {
	case "SQL":
		dbStorage := sqlstorage.NewDBStorage()
		if err := dbStorage.Connect(ctx, dsn); err != nil {
			return nil, fmt.Errorf("failed to init db storage: %w", err)
		}
		db = dbStorage

		/*defer func() {
			if err := dbStorage.Close(ctx); err != nil {
				return
			}
		}()*/
	default:
		storageMemory := memorystorage.New()
		db = storageMemory
	}
	return db, nil
}
