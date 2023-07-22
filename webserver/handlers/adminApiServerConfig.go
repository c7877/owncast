package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/services/config"
	"github.com/owncast/owncast/utils"
	"github.com/owncast/owncast/video/transcoder"
	"github.com/owncast/owncast/webserver/middleware"
	log "github.com/sirupsen/logrus"
)

// GetServerConfig gets the config details of the server.
func (h *Handlers) GetServerConfig(w http.ResponseWriter, r *http.Request) {
	ffmpeg := utils.ValidatedFfmpegPath(configRepository.GetFfMpegPath())
	usernameBlocklist := configRepository.GetForbiddenUsernameList()
	usernameSuggestions := configRepository.GetSuggestedUsernamesList()
	c := config.Get()

	videoQualityVariants := make([]models.StreamOutputVariant, 0)
	for _, variant := range configRepository.GetStreamOutputVariants() {
		videoQualityVariants = append(videoQualityVariants, models.StreamOutputVariant{
			Name:               variant.GetName(),
			IsAudioPassthrough: variant.GetIsAudioPassthrough(),
			IsVideoPassthrough: variant.IsVideoPassthrough,
			Framerate:          variant.GetFramerate(),
			VideoBitrate:       variant.VideoBitrate,
			AudioBitrate:       variant.AudioBitrate,
			CPUUsageLevel:      variant.CPUUsageLevel,
			ScaledWidth:        variant.ScaledWidth,
			ScaledHeight:       variant.ScaledHeight,
		})
	}
	response := serverConfigAdminResponse{
		InstanceDetails: adminWebConfigResponse{
			Name:                configRepository.GetServerName(),
			Summary:             configRepository.GetServerSummary(),
			Tags:                configRepository.GetServerMetadataTags(),
			ExtraPageContent:    configRepository.GetExtraPageBodyContent(),
			StreamTitle:         configRepository.GetStreamTitle(),
			WelcomeMessage:      configRepository.GetServerWelcomeMessage(),
			OfflineMessage:      configRepository.GetCustomOfflineMessage(),
			Logo:                configRepository.GetLogoPath(),
			SocialHandles:       configRepository.GetSocialHandles(),
			NSFW:                configRepository.GetNSFW(),
			CustomStyles:        configRepository.GetCustomStyles(),
			CustomJavascript:    configRepository.GetCustomJavascript(),
			AppearanceVariables: configRepository.GetCustomColorVariableValues(),
		},
		FFmpegPath:              ffmpeg,
		AdminPassword:           configRepository.GetAdminPassword(),
		StreamKeys:              configRepository.GetStreamKeys(),
		StreamKeyOverridden:     c.TemporaryStreamKey != "",
		WebServerPort:           c.WebServerPort,
		WebServerIP:             c.WebServerIP,
		RTMPServerPort:          configRepository.GetRTMPPortNumber(),
		ChatDisabled:            configRepository.GetChatDisabled(),
		ChatJoinMessagesEnabled: configRepository.GetChatJoinMessagesEnabled(),
		SocketHostOverride:      configRepository.GetWebsocketOverrideHost(),
		VideoServingEndpoint:    configRepository.GetVideoServingEndpoint(),
		ChatEstablishedUserMode: configRepository.GetChatEstbalishedUsersOnlyMode(),
		HideViewerCount:         configRepository.GetHideViewerCount(),
		DisableSearchIndexing:   configRepository.GetDisableSearchIndexing(),
		VideoSettings: videoSettings{
			VideoQualityVariants: videoQualityVariants,
			LatencyLevel:         configRepository.GetStreamLatencyLevel().Level,
		},
		YP: yp{
			Enabled:     configRepository.GetDirectoryEnabled(),
			InstanceURL: configRepository.GetServerURL(),
		},
		S3:                 configRepository.GetS3Config(),
		ExternalActions:    configRepository.GetExternalActions(),
		SupportedCodecs:    transcoder.GetCodecs(ffmpeg),
		VideoCodec:         configRepository.GetVideoCodec(),
		ForbiddenUsernames: usernameBlocklist,
		SuggestedUsernames: usernameSuggestions,
		Federation: adminFederationConfigResponse{
			Enabled:        configRepository.GetFederationEnabled(),
			IsPrivate:      configRepository.GetFederationIsPrivate(),
			Username:       configRepository.GetFederationUsername(),
			GoLiveMessage:  configRepository.GetFederationGoLiveMessage(),
			ShowEngagement: configRepository.GetFederationShowEngagement(),
			BlockedDomains: configRepository.GetBlockedFederatedDomains(),
		},
		Notifications: adminNotificationsConfigResponse{
			Discord: configRepository.GetDiscordConfig(),
			Browser: configRepository.GetBrowserPushConfig(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	middleware.DisableCache(w)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorln(err)
	}
}

type serverConfigAdminResponse struct {
	InstanceDetails         adminWebConfigResponse           `json:"instanceDetails"`
	Notifications           adminNotificationsConfigResponse `json:"notifications"`
	YP                      yp                               `json:"yp"`
	FFmpegPath              string                           `json:"ffmpegPath"`
	AdminPassword           string                           `json:"adminPassword"`
	SocketHostOverride      string                           `json:"socketHostOverride,omitempty"`
	WebServerIP             string                           `json:"webServerIP"`
	VideoCodec              string                           `json:"videoCodec"`
	VideoServingEndpoint    string                           `json:"videoServingEndpoint"`
	S3                      models.S3                        `json:"s3"`
	Federation              adminFederationConfigResponse    `json:"federation"`
	SupportedCodecs         []string                         `json:"supportedCodecs"`
	ExternalActions         []models.ExternalAction          `json:"externalActions"`
	ForbiddenUsernames      []string                         `json:"forbiddenUsernames"`
	SuggestedUsernames      []string                         `json:"suggestedUsernames"`
	StreamKeys              []models.StreamKey               `json:"streamKeys"`
	VideoSettings           videoSettings                    `json:"videoSettings"`
	RTMPServerPort          int                              `json:"rtmpServerPort"`
	WebServerPort           int                              `json:"webServerPort"`
	ChatDisabled            bool                             `json:"chatDisabled"`
	ChatJoinMessagesEnabled bool                             `json:"chatJoinMessagesEnabled"`
	ChatEstablishedUserMode bool                             `json:"chatEstablishedUserMode"`
	DisableSearchIndexing   bool                             `json:"disableSearchIndexing"`
	StreamKeyOverridden     bool                             `json:"streamKeyOverridden"`
	HideViewerCount         bool                             `json:"hideViewerCount"`
}

type videoSettings struct {
	VideoQualityVariants []models.StreamOutputVariant `json:"videoQualityVariants"`
	LatencyLevel         int                          `json:"latencyLevel"`
}

type adminWebConfigResponse struct {
	AppearanceVariables map[string]string     `json:"appearanceVariables"`
	Version             string                `json:"version"`
	WelcomeMessage      string                `json:"welcomeMessage"`
	OfflineMessage      string                `json:"offlineMessage"`
	Logo                string                `json:"logo"`
	Name                string                `json:"name"`
	ExtraPageContent    string                `json:"extraPageContent"`
	StreamTitle         string                `json:"streamTitle"` // What's going on with the current stream
	CustomStyles        string                `json:"customStyles"`
	CustomJavascript    string                `json:"customJavascript"`
	Summary             string                `json:"summary"`
	Tags                []string              `json:"tags"`
	SocialHandles       []models.SocialHandle `json:"socialHandles"`
	NSFW                bool                  `json:"nsfw"`
}

type yp struct {
	InstanceURL  string `json:"instanceUrl"` // The public URL the directory should link to
	YPServiceURL string `json:"-"`           // The base URL to the YP API to register with (optional)
	Enabled      bool   `json:"enabled"`
}

type adminFederationConfigResponse struct {
	Username       string   `json:"username"`
	GoLiveMessage  string   `json:"goLiveMessage"`
	BlockedDomains []string `json:"blockedDomains"`
	Enabled        bool     `json:"enabled"`
	IsPrivate      bool     `json:"isPrivate"`
	ShowEngagement bool     `json:"showEngagement"`
}

type adminNotificationsConfigResponse struct {
	Browser models.BrowserNotificationConfiguration `json:"browser"`
	Discord models.DiscordConfiguration             `json:"discord"`
}
