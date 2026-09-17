package errors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/infrastructure/logger"

	"gorm.io/gorm"
)

const pgUniqueViolation = "23505"

// Wrap turns a raw persistence error into a *response.Response the API
// layer knows how to render, logging anything unexpected along the way.
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFoundError
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation {
		return response.ConflictError
	}

	logger.InternalLogger.Error(err.Error())
	return err
}
