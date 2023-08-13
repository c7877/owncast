package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/owncast/owncast/activitypub"
	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/services/config"
	"github.com/owncast/owncast/utils"
	"github.com/owncast/owncast/webserver/middleware"
	"github.com/owncast/owncast/webserver/responses"
	log "github.com/sirupsen/logrus"
)

type webConfigResponse struct {
	AppearanceVariables  map[string]string            `json:"appearanceVariables"`
	Notifications        notificationsConfigResponse  `json:"notifications"`
	CustomStyles         string                       `json:"customStyles"`
	Summary              string                       `json:"summary"`
	OfflineMessage       string                       `json:"offlineMessage"`
	Logo                 string                       `json:"logo"`
	Version              string                       `json:"version"`
	SocketHostOverride   string                       `json:"socketHostOverride,omitempty"`
	ExtraPageContent     string                       `json:"extraPageContent"`
	StreamTitle          string                       `json:"streamTitle,omitempty"` // What's going on with the current stream
	Name                 string                       `json:"name"`
	Federation           federationConfigResponse     `json:"federation"`
	SocialHandles        []models.SocialHandle        `json:"socialHandles"`
	ExternalActions      []models.ExternalAction      `json:"externalActions"`
	Tags                 []string                     `json:"tags"`
	MaxSocketPayloadSize int                          `json:"maxSocketPayloadSize"`
	ChatDisabled         bool                         `json:"chatDisabled"`
	NSFW                 bool                         `json:"nsfw"`
	Authentication       authenticationConfigResponse `json:"authentication"`
	HideViewerCount      bool                         `json:"hideViewerCount"`
}

type federationConfigResponse struct {
	Account       string `json:"account,omitempty"`
	FollowerCount int    `json:"followerCount,omitempty"`
	Enabled       bool   `json:"enabled"`
}

type browserNotificationsConfigResponse struct {
	PublicKey string `json:"publicKey,omitempty"`
	Enabled   bool   `json:"enabled"`
}

type notificationsConfigResponse struct {
	Browser browserNotificationsConfigResponse `json:"browser"`
}

type authenticationConfigResponse struct {
	IndieAuthEnabled bool `json:"indieAuthEnabled"`
}

// GetWebConfig gets the status of the server.
func (h *Handlers) GetWebConfig(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(w)
	middleware.DisableCache(w)
	w.Header().Set("Content-Type", "application/json")

	configuration := getConfigResponse()

	if err := json.NewEncoder(w).Encode(configuration); err != nil {
		responses.BadRequestHandler(w, err)
	}
}

func getConfigResponse() webConfigResponse {
	pageContent := utils.RenderPageContentMarkdown(configRepository.GetExtraPageBodyContent())
	offlineMessage := utils.RenderSimpleMarkdown(configRepository.GetCustomOfflineMessage())
	socialHandles := configRepository.GetSocialHandles()
	for i, handle := range socialHandles {
		platform := models.GetSocialHandle(handle.Platform)
		if platform != nil {
			handle.Icon = platform.Icon
			socialHandles[i] = handle
		}
	}

	serverSummary := configRepository.GetServerSummary()

	var federationResponse federationConfigResponse
	federationEnabled := configRepository.GetFederationEnabled()

	followerCount, _ := activitypub.GetFollowerCount()
	if federationEnabled {
		serverURLString := configRepository.GetServerURL()
		serverURL, _ := url.Parse(serverURLString)
		account := fmt.Sprintf("%s@%s", configRepository.GetDefaultFederationUsername(), serverURL.Host)
		federationResponse = federationConfigResponse{
			Enabled:       federationEnabled,
			FollowerCount: int(followerCount),
			Account:       account,
		}
	}

	browserPushEnabled := configRepository.GetBrowserPushConfig().Enabled
	browserPushPublicKey, err := configRepository.GetBrowserPushPublicKey()
	if err != nil {
		log.Errorln("unable to fetch browser push notifications public key", err)
		browserPushEnabled = false
	}

	notificationsResponse := notificationsConfigResponse{
		Browser: browserNotificationsConfigResponse{
			Enabled:   browserPushEnabled,
			PublicKey: browserPushPublicKey,
		},
	}

	authenticationResponse := authenticationConfigResponse{
		IndieAuthEnabled: configRepository.GetServerURL() != "",
	}

	c := config.GetConfig()

	return webConfigResponse{
		Name:                 configRepository.GetServerName(),
		Summary:              serverSummary,
		OfflineMessage:       offlineMessage,
		Logo:                 "/logo",
		Tags:                 configRepository.GetServerMetadataTags(),
		Version:              c.GetReleaseString(),
		NSFW:                 configRepository.GetNSFW(),
		SocketHostOverride:   configRepository.GetWebsocketOverrideHost(),
		ExtraPageContent:     pageContent,
		StreamTitle:          configRepository.GetStreamTitle(),
		SocialHandles:        socialHandles,
		ChatDisabled:         configRepository.GetChatDisabled(),
		ExternalActions:      configRepository.GetExternalActions(),
		CustomStyles:         configRepository.GetCustomStyles(),
		MaxSocketPayloadSize: config.MaxSocketPayloadSize,
		Federation:           federationResponse,
		Notifications:        notificationsResponse,
		Authentication:       authenticationResponse,
		AppearanceVariables:  configRepository.GetCustomColorVariableValues(),
		HideViewerCount:      configRepository.GetHideViewerCount(),
	}
}

// GetAllSocialPlatforms will return a list of all social platform types.
func (h *Handlers) GetAllSocialPlatforms(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(w)
	w.Header().Set("Content-Type", "application/json")

	platforms := models.GetAllSocialHandles()
	if err := json.NewEncoder(w).Encode(platforms); err != nil {
		responses.InternalErrorHandler(w, err)
	}
}
