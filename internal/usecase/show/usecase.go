package show

import "database/sql"

type UseCase struct {
	db       *sql.DB
	showRepo showRepository
	seatRepo seatRepository
}

func New(
	db *sql.DB,
	showRepo showRepository,
	seatRepo seatRepository,
) *UseCase {
	return &UseCase{
		db:       db,
		showRepo: showRepo,
		seatRepo: seatRepo,
	}
}
