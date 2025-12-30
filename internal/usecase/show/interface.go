package show

import (
	"context"
	"database/sql"
)

type showRepository interface {
	ShowExistsTx(context.Context, *sql.Tx, int) (bool, error)
}

type seatRepository interface {
	CheckAndReserveTx(context.Context, *sql.Tx, int, []int) error
}
