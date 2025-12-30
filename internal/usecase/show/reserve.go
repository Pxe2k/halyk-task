package show

import (
	"context"
	"database/sql"
	"time"

	"github.com/Pxe2k/halyk-task/internal/delivery/show"
	"github.com/Pxe2k/halyk-task/internal/domain"
	"github.com/Pxe2k/halyk-task/pkg"
)

const (
	maxRetryAttempts = 3
	retryDelay       = 50 * time.Millisecond
)

func (uc UseCase) ReserveSeats(ctx context.Context, req show.ReserveSeatsRequest) error {
	var lastErr error

	for attempt := 1; attempt <= maxRetryAttempts; attempt++ {
		err := uc.reserveSeatsWithTx(ctx, req)

		if err == nil {
			return nil
		}

		if err != domain.ErrDeadlock || attempt == maxRetryAttempts {
			return err
		}

		lastErr = err

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(uc.calculateBackoff(attempt)):
		}
	}

	return lastErr
}

func (uc UseCase) reserveSeatsWithTx(ctx context.Context, req show.ReserveSeatsRequest) error {
	tx, err := uc.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	exists, err := uc.showRepo.ShowExistsTx(ctx, tx, req.ShowID)
	if err != nil {
		return err
	}
	if !exists {
		return domain.ErrShowNotFound
	}

	if err = uc.seatRepo.CheckAndReserveTx(ctx, tx, req.ShowID, req.SeatNumbers); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		if pkg.IsDeadlock(err) {
			return domain.ErrDeadlock
		}
		return err
	}

	return nil
}

func (uc UseCase) calculateBackoff(attempt int) time.Duration {
	backoff := retryDelay * time.Duration(1<<(attempt-1))

	return backoff
}
