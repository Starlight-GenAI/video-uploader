package core

import (
	"github.com/dreammnck/video-uploader/pkg/v1/model"
	"google.golang.org/api/youtube/v3"
)

type videoUploaderSvc struct {
	pubsubAddapter           model.IPubSubAdapter
	youtubeService           *youtube.Service
	queueHistoryFirebaseRepo model.IQueueHistoryFirestoreRepo
}

func NewVideoUploaderSvc(pubsubAddapter model.IPubSubAdapter, youtubeService *youtube.Service, queueHistoryFirebaseRepo model.IQueueHistoryFirestoreRepo) *videoUploaderSvc {
	return &videoUploaderSvc{
		pubsubAddapter:           pubsubAddapter,
		youtubeService:           youtubeService,
		queueHistoryFirebaseRepo: queueHistoryFirebaseRepo,
	}
}

var _ model.IVideoUplaoderSvc = new(videoUploaderSvc)
