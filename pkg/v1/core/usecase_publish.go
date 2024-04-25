package core

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dreammnck/video-uploader/pkg/logger"
	"github.com/dreammnck/video-uploader/pkg/v1/model"
	"github.com/google/uuid"
	"github.com/sagikazarmark/slog-shim"
)

const pending = "pending"

func (s *videoUploaderSvc) Publish(ctx context.Context, opts model.PublishOpts) (*string, error) {

	logger.Logger.Info("publish event", slog.String("tag", "usecase publish event"))
	id := strings.ReplaceAll(uuid.NewString(), "-", "")
	videoUrl := ""
	videoID := ""

	if opts.VideoUrl == "" {
		videoUrl = fmt.Sprintf(model.YoutubeLinkeScheme, opts.VideoID)
		videoID = opts.VideoID
	} else {
		videoUrl = opts.VideoUrl
		videoIDPtr, err := s.getVideoID(videoUrl)
		if err != nil {
			return nil, err
		}
		videoID = *videoIDPtr
	}

	videoDetail, err := s.GetVideoDetail(ctx, videoUrl)
	if err != nil {
		return nil, err
	}

	if err := s.pubsubAddapter.Publish(ctx, model.VideoExtractEventMessage{
		ID:            id,
		VideoUrl:      videoUrl,
		IsUseSubTitle: opts.IsUseSubTitle,
		UserID:        opts.UserID,
	}); err != nil {
		logger.Logger.Error(fmt.Sprintf("publish event fail with %s", err.Error()), slog.String("tag", "usecase upload video"))
		return nil, err
	}

	now := time.Now()

	if err := s.queueHistoryFirebaseRepo.Insert(ctx, model.QueueHistory{
		ID:            id,
		ChannelName:   videoDetail.ChannelName,
		UserID:        opts.UserID,
		VideoUrl:      videoUrl,
		VideoID:       videoID,
		Title:         videoDetail.Title,
		Description:   videoDetail.Description,
		Thumbnails:    videoDetail.Thumbnails,
		Status:        pending,
		CreatedAt:     now,
		UpdatedAt:     now,
		IsUseSubTitle: opts.IsUseSubTitle,
	}); err != nil {
		return nil, err
	}

	logger.Logger.Info("successfully upload video", slog.String("tag", "usecase publish event"))
	return &id, nil
}
