package core

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/dreammnck/video-uploader/pkg/logger"
	"github.com/dreammnck/video-uploader/pkg/v1/model"
	"github.com/sagikazarmark/slog-shim"
)

func (s *videoUploaderSvc) GetVideoDetail(ctx context.Context, video_url string) (*model.VideoDetail, error) {
	logger.Logger.Info("get video detail", slog.String("tag", "usecase get video detail"))

	video_id, err := s.getVideoID(video_url)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("extract video id fail with %s", err.Error()), slog.String("tag", "usecase get video detail"))
		return nil, err
	}

	videoResponse, err := s.youtubeService.Videos.
		List([]string{"statistics", "snippet", "contentDetails"}).
		Id(*video_id).
		Do()

	if err != nil {
		logger.Logger.Error(fmt.Sprintf("call youtube service fail with %s", err.Error()), slog.String("tag", "usecase get video detail"))
		return nil, err
	}

	if len(videoResponse.Items) == 0 {
		return nil, nil
	}

	video := videoResponse.Items[0]
	channel, err := s.getChannelDetail(ctx, video.Snippet.ChannelId)
	if err != nil {
		return nil, err
	}

	logger.Logger.Info("successfully get video detail", slog.String("tag", "usecase get video detail"))
	return &model.VideoDetail{
		ChannelName: channel.ChannelName,
		Title:       video.Snippet.Title,
		Description: video.Snippet.Description,
		Thumbnails:  video.Snippet.Thumbnails.Medium.Url,
		PublishAt:   video.Snippet.PublishedAt,
		Duration:    video.ContentDetails.Duration,
		ViewCount:   video.Statistics.ViewCount,
		LikeCount:   video.Statistics.LikeCount,
	}, nil
}

func (s *videoUploaderSvc) getVideoID(video_url string) (*string, error) {
	regexPattern := `(?:youtube\.com\/(?:[^\/]+\/.+\/|(?:v|e(?:mbed)?)\/|.*[?&]v=)|youtu\.be\/)([^"&?\/\s]{11})`
	regex := regexp.MustCompile(regexPattern)
	matches := regex.FindStringSubmatch(video_url)
	if len(matches) < 2 {
		return nil, errors.New("invalid youtube url")
	}

	return &matches[1], nil
}

func (s *videoUploaderSvc) getChannelDetail(ctx context.Context, channelID string) (*model.ChannelDetail, error) {

	channelResponse, err := s.youtubeService.Channels.List([]string{"snippet"}).
		Id(channelID).
		Do()
	if err != nil {
		return nil, err
	}

	if len(channelResponse.Items) == 0 {
		return nil, nil
	}
	channel := channelResponse.Items[0]

	return &model.ChannelDetail{
		ChannelName: channel.Snippet.Title,
	}, nil
}
