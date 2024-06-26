package resthendler

import (
	"encoding/json"
	"net/http"

	celenderdomain "main.go/Downloads/L2/develop/dev11/internal/domain"
	"main.go/Downloads/L2/develop/dev11/internal/service/celender"
	restutills "main.go/Downloads/L2/develop/dev11/internal/transport/rest/utills"

	"github.com/google/uuid"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if restutills.ValidateQuery(w, r, http.MethodPost) {
		event := celenderdomain.Event{}
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			restutills.SendError(w, http.StatusServiceUnavailable, err.Error())
		} else {
			event.Uid = uuid.New().String()
			celender.Instanse.Add(event.User, celender.NewEventByDomain(event))
			restutills.SendResponce(w, http.StatusOK, "created")
		}
	}

}
