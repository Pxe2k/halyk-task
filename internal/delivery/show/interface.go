package show

import (
	"context"
)

type UseCase interface {
	ReserveSeats(ctx context.Context, req ReserveSeatsRequest) error
}
