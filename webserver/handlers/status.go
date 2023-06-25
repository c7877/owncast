package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/owncast/owncast/core"
	"github.com/owncast/owncast/utils"
	"github.com/owncast/owncast/webserver/middleware"
	"github.com/owncast/owncast/webserver/responses"
)

// GetStatus gets the status of the server.
func (h *Handlers) GetStatus(w http.ResponseWriter, r *http.Request) {
	response := getStatusResponse()

	w.Header().Set("Content-Type", "application/json")
	middleware.DisableCache(w)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		responses.InternalErrorHandler(w, err)
	}
}

func getStatusResponse() webStatusResponse {
	status := core.GetStatus()
	response := webStatusResponse{
		Online:             status.Online,
		ServerTime:         time.Now(),
		LastConnectTime:    status.LastConnectTime,
		LastDisconnectTime: status.LastDisconnectTime,
		VersionNumber:      status.VersionNumber,
		StreamTitle:        status.StreamTitle,
	}
	if !configRepository.GetHideViewerCount() {
		response.ViewerCount = status.ViewerCount
	}
	return response
}

type webStatusResponse struct {
	ServerTime         time.Time       `json:"serverTime"`
	LastConnectTime    *utils.NullTime `json:"lastConnectTime"`
	LastDisconnectTime *utils.NullTime `json:"lastDisconnectTime"`

	VersionNumber string `json:"versionNumber"`
	StreamTitle   string `json:"streamTitle"`
	ViewerCount   int    `json:"viewerCount,omitempty"`
	Online        bool   `json:"online"`
}
