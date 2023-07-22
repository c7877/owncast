package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/webserver/responses"
	log "github.com/sirupsen/logrus"
)

// GetUserDetails returns the details of a chat user for moderators.
func (h *Handlers) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	type connectedClient struct {
		ConnectedAt  time.Time `json:"connectedAt"`
		UserAgent    string    `json:"userAgent"`
		Geo          string    `json:"geo,omitempty"`
		Id           uint      `json:"id"`
		MessageCount int       `json:"messageCount"`
	}

	type response struct {
		User             *models.User              `json:"user"`
		ConnectedClients []connectedClient         `json:"connectedClients"`
		Messages         []models.UserMessageEvent `json:"messages"`
	}

	pathComponents := strings.Split(r.URL.Path, "/")
	uid := pathComponents[len(pathComponents)-1]

	u := userRepository.GetUserByID(uid)

	if u == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	c, _ := h.chatService.GetClientsForUser(uid)
	clients := make([]connectedClient, len(c))
	for i, c := range c {
		client := connectedClient{
			Id:           c.Id,
			MessageCount: c.MessageCount,
			UserAgent:    c.UserAgent,
			ConnectedAt:  c.ConnectedAt,
		}
		if c.Geo != nil {
			client.Geo = c.Geo.CountryCode
		}

		clients[i] = client
	}

	messages, err := h.chatRepository.GetMessagesFromUser(uid)
	if err != nil {
		log.Errorln(err)
	}

	res := response{
		User:             u,
		ConnectedClients: clients,
		Messages:         messages,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		responses.InternalErrorHandler(w, err)
	}
}
