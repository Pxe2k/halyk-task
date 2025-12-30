package show

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Pxe2k/halyk-task/internal/delivery/show"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockShowRepo struct {
	mock.Mock
}

func (m *mockShowRepo) ShowExistsTx(ctx context.Context, tx *sql.Tx, showID int) (bool, error) {
	args := m.Called(ctx, tx, showID)
	return args.Bool(0), args.Error(1)
}

type mockSeatRepo struct {
	mock.Mock
}

func (m *mockSeatRepo) CheckAndReserveTx(ctx context.Context, tx *sql.Tx, showID int, seatNumbers []int) error {
	args := m.Called(ctx, tx, showID, seatNumbers)
	return args.Error(0)
}

func TestReserveSeats_Success(t *testing.T) {
	ctx := context.Background()
	req := show.ReserveSeatsRequest{
		ShowID:      1,
		SeatNumbers: []int{1, 2, 3},
	}

	db, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	showRepo := &mockShowRepo{}
	seatRepo := &mockSeatRepo{}

	mockDB.ExpectBegin()

	uc := UseCase{
		db:       db,
		showRepo: showRepo,
		seatRepo: seatRepo,
	}

	showRepo.On("ShowExistsTx", ctx, mock.AnythingOfType("*sql.Tx"), 1).Return(true, nil)
	seatRepo.On("CheckAndReserveTx", ctx, mock.AnythingOfType("*sql.Tx"), 1, []int{1, 2, 3}).Return(nil)
	mockDB.ExpectCommit()

	err = uc.ReserveSeats(ctx, req)
	assert.NoError(t, err)

	assert.NoError(t, mockDB.ExpectationsWereMet())
	showRepo.AssertExpectations(t)
	seatRepo.AssertExpectations(t)
}

func TestReserveSeats_PartialAvailability(t *testing.T) {
	ctx := context.Background()
	req := show.ReserveSeatsRequest{
		ShowID:      1,
		SeatNumbers: []int{1, 2, 3},
	}

	db, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	showRepo := &mockShowRepo{}
	seatRepo := &mockSeatRepo{}

	mockDB.ExpectBegin()

	uc := UseCase{
		db:       db,
		showRepo: showRepo,
		seatRepo: seatRepo,
	}

	showRepo.On("ShowExistsTx", ctx, mock.AnythingOfType("*sql.Tx"), 1).Return(true, nil)
	seatRepo.On("CheckAndReserveTx", ctx, mock.AnythingOfType("*sql.Tx"), 1, []int{1, 2, 3}).Return(sql.ErrNoRows)
	mockDB.ExpectRollback()

	err = uc.ReserveSeats(ctx, req)
	assert.Error(t, err)

	assert.NoError(t, mockDB.ExpectationsWereMet())
	showRepo.AssertExpectations(t)
	seatRepo.AssertExpectations(t)
}

func TestReserveSeats_ShowNotFound(t *testing.T) {
	ctx := context.Background()
	req := show.ReserveSeatsRequest{
		ShowID:      999,
		SeatNumbers: []int{1},
	}

	db, mockDB, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	showRepo := &mockShowRepo{}
	seatRepo := &mockSeatRepo{}

	mockDB.ExpectBegin()

	uc := UseCase{
		db:       db,
		showRepo: showRepo,
		seatRepo: seatRepo,
	}

	showRepo.On("ShowExistsTx", ctx, mock.AnythingOfType("*sql.Tx"), 999).Return(false, nil)
	mockDB.ExpectRollback()

	err = uc.ReserveSeats(ctx, req)
	assert.Error(t, err)

	assert.NoError(t, mockDB.ExpectationsWereMet())
	showRepo.AssertExpectations(t)
	seatRepo.AssertNotCalled(t, "CheckAndReserveTx")
}

func TestReserveSeats_Concurrent(t *testing.T) {
	t.Parallel()

	done := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(id int) {
			ctx := context.Background()
			req := show.ReserveSeatsRequest{
				ShowID:      id,
				SeatNumbers: []int{id},
			}

			db, mockDB, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			showRepo := &mockShowRepo{}
			seatRepo := &mockSeatRepo{}

			mockDB.ExpectBegin()
			showRepo.On("ShowExistsTx", ctx, mock.AnythingOfType("*sql.Tx"), id).Return(true, nil)
			seatRepo.On("CheckAndReserveTx", ctx, mock.AnythingOfType("*sql.Tx"), id, []int{id}).Return(nil)
			mockDB.ExpectCommit()

			uc := UseCase{
				db:       db,
				showRepo: showRepo,
				seatRepo: seatRepo,
			}

			err = uc.ReserveSeats(ctx, req)
			assert.NoError(t, err)

			done <- true
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-done
	}
}
