package show

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Pxe2k/halyk-task/pkg"
)

func (h *Handler) researveSeats(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	var in ReserveSeatsRequest
	if err = json.Unmarshal(body, &in); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	if err = in.validate(); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	if err = h.uc.ReserveSeats(r.Context(), in); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.ToJSON(w, http.StatusOK, map[string]string{
		"status": "success",
	})
}
