package show

import (
	"github.com/Pxe2k/halyk-task/internal/domain"
)

type (
	ReserveSeatsRequest struct {
		ShowID      int   `json:"show_id"`
		SeatNumbers []int `json:"seat_numbers"`
	}
)

func (r *ReserveSeatsRequest) validate() error {
	if r.ShowID <= 0 {
		return domain.ErrShowIDValidate
	}
	if len(r.SeatNumbers) == 0 {
		return domain.ErrSeatNumbersValidate
	}

	return nil
}
