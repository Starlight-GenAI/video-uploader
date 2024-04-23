package routes

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/video-uploader/config"
	"github.com/dreammnck/video-uploader/pkg/v1/adapter"
	"github.com/dreammnck/video-uploader/pkg/v1/core"
	"github.com/dreammnck/video-uploader/pkg/v1/handler"
	"github.com/dreammnck/video-uploader/pkg/v1/repo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/youtube/v3"
)

type apiRouter struct {
	config          *config.Config
	pubSubClient    *pubsub.Client
	youtubeService  *youtube.Service
	firestoreClient *firestore.Client
}

func NewRouter(config *config.Config, pubSubClient *pubsub.Client, youtubeService *youtube.Service, firestoreClient *firestore.Client) *apiRouter {
	return &apiRouter{config: config, pubSubClient: pubSubClient, youtubeService: youtubeService, firestoreClient: firestoreClient}
}

func (a *apiRouter) RegisterRouter() *echo.Echo {

	router := newEcho()
	pubsubAdapter := adapter.NewPubSubAdapter(a.pubSubClient, a.config.PubSub.Topic)
	queueHistoryFirebaseRepo := repo.NewQueueHistoryFirestoreRepo(a.firestoreClient, a.config.Firestore.QueueHistoryCollection)
	videoUploaderSvc := core.NewVideoUploaderSvc(pubsubAdapter, a.youtubeService, queueHistoryFirebaseRepo)
	h := handler.NewVideoUploaderHandler(videoUploaderSvc)

	router.POST("/upload", h.UploadVideoHandler)
	router.POST("/video-info", h.VideoInfo)
	router.GET("/health", func(e echo.Context) error {
		return e.String(http.StatusOK, "OK")
	})

	return router
}

func newEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Secure())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
		HSTSExcludeSubdomains: true,
	}))
	e.Use(middleware.CORS())

	return e
}
