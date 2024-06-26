package resthendler

import (
	"encoding/json"
	"net/http"

	celenderdomain "main.go/Downloads/L2/develop/dev11/internal/domain"
	"main.go/Downloads/L2/develop/dev11/internal/service/celender"
	restutills "main.go/Downloads/L2/develop/dev11/internal/transport/rest/utills"
)

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if restutills.ValidateQuery(w, r, http.MethodPost) {
		event := celenderdomain.Event{}
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil || !event.Validate() {
			restutills.SendError(w, http.StatusServiceUnavailable, "not parameter UID or isn't correct")
			return
		}
		if ok := celender.Instanse.Del(event.User, event.Uid); !ok {
			restutills.SendError(w, http.StatusNotModified, "not found")
			return
		}
		restutills.SendResponce(w, http.StatusOK, "deleted")
	}
}
