package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/video-uploader/config"
	routes "github.com/dreammnck/video-uploader/pkg"
	"github.com/dreammnck/video-uploader/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/sagikazarmark/slog-shim"
	"go.uber.org/zap"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {

	configs := config.InitConfig()
	ctx := context.Background()
	logger.Init()

	pubSubClient, err := pubsub.NewClient(ctx, configs.PubSub.ProjectID, option.WithCredentialsFile(configs.PubSub.CredentialPath))
	if err != nil {
		panic(err)
	}

	defer pubSubClient.Close()

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey("AIzaSyAB5WD38r8E9qcR6M793kZ9yiqEZCowFOU"))
	if err != nil {
		panic(err)
	}

	fireStoreClient, err := firestore.NewClientWithDatabase(ctx, configs.Firestore.ProjectID, configs.Firestore.Database, option.WithCredentialsFile(configs.Firestore.CredentialFilePath))
	if err != nil {
		panic(err)
	}
	defer fireStoreClient.Close()

	apiRouter := routes.NewRouter(configs, pubSubClient, youtubeService, fireStoreClient).RegisterRouter()
	go RunServer(apiRouter, configs.Server.Port)

	Shutdown(apiRouter)
}

func RunServer(router *echo.Echo, port int) {
	startPort := fmt.Sprintf(":%d", port)
	router.Logger.Fatal(router.Start(startPort))
}

func Shutdown(router *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	if err := router.Shutdown(context.Background()); err != nil {
		slog.Error(err.Error(), zap.String("tag", "shutdown Server"))
	}
}
