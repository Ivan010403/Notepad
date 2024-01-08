package delete

import (
	"net/http"
	"notepad/internal/storage/postgres"

	"go.uber.org/zap"
)

func Delete(logger *zap.Logger, db *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
