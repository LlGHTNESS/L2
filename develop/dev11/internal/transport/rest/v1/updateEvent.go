package resthendler

import (
	"encoding/json"
	"net/http"

	restutills "main.go/Downloads/L2/develop/dev11/internal/transport/rest/utills"

	celenderdomain "main.go/Downloads/L2/develop/dev11/internal/domain"
	"main.go/Downloads/L2/develop/dev11/internal/service/celender"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if restutills.ValidateQuery(w, r, http.MethodPost) {
		event := celenderdomain.Event{}
		if ok := json.NewDecoder(r.Body).Decode(&event); ok != nil {
			restutills.SendError(w, http.StatusServiceUnavailable, ok.Error())
		} else if ok := celender.Instanse.Update(event.User, celender.NewEventByDomain(event)); ok != nil {
			restutills.SendError(w, http.StatusServiceUnavailable, ok.Error())
		} else {
			restutills.SendResponce(w, http.StatusOK, "updated")
		}
	}
}
