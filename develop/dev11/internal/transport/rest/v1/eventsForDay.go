package resthendler

import (
	"net/http"
	"strconv"
	"time"

	"main.go/Downloads/L2/develop/dev11/internal/service/celender"
	restutills "main.go/Downloads/L2/develop/dev11/internal/transport/rest/utills"
)

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	if restutills.ValidateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, ok := time.Parse("2006-01-02", r.URL.Query().Get("date")); ok != nil {
			restutills.SendError(w, http.StatusServiceUnavailable, ok.Error())
		} else {
			userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
			if err != nil {
				restutills.SendError(w, http.StatusServiceUnavailable, err.Error())
			}
			evs := celender.Instanse.GetFromTimeInterval(userID, date, date.AddDate(0, 0, 1))
			restutills.SendResponce(w, http.StatusOK, evs)
		}
	}
}
