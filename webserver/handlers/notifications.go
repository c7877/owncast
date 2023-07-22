package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/services/notifications"
	"github.com/owncast/owncast/storage/notificationsrepository"
	"github.com/owncast/owncast/webserver/responses"

	"github.com/owncast/owncast/utils"

	log "github.com/sirupsen/logrus"
)

// RegisterForLiveNotifications will register a channel + destination to be
// notified when a stream goes live.
func (h *Handlers) RegisterForLiveNotifications(u models.User, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responses.WriteSimpleResponse(w, false, r.Method+" not supported")
		return
	}

	type request struct {
		// Channel is the notification channel (browser, sms, etc)
		Channel string `json:"channel"`
		// Destination is the target of the notification in the above channel.
		Destination string `json:"destination"`
	}

	decoder := json.NewDecoder(r.Body)
	var req request
	if err := decoder.Decode(&req); err != nil {
		log.Errorln(err)
		responses.WriteSimpleResponse(w, false, "unable to register for notifications")
		return
	}

	// Make sure the requested channel is one we want to handle.
	validTypes := []string{notifications.BrowserPushNotification}
	_, validChannel := utils.FindInSlice(validTypes, req.Channel)
	if !validChannel {
		responses.WriteSimpleResponse(w, false, "invalid notification channel: "+req.Channel)
		return
	}

	n := notificationsrepository.Get()

	if err := n.AddNotification(req.Channel, req.Destination); err != nil {
		log.Errorln(err)
		responses.WriteSimpleResponse(w, false, "unable to save notification")
		return
	}
}
