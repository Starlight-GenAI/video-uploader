package model

import (
	"context"
)

type IVideoUplaoderSvc interface {
	Publish(ctx context.Context, opts PublishOpts) (*string, error)
	GetVideoDetail(ctx context.Context, video_url string) (*VideoDetail, error)
}

type IPubSubAdapter interface {
	Publish(ctx context.Context, data VideoExtractEventMessage) error
}

type IQueueHistoryFirestoreRepo interface {
	Insert(ctx context.Context, data QueueHistory) error
}

type PublishOpts struct {
	VideoID       string
	UserID        string
	VideoUrl      string
	IsUseSubTitle bool
}
type VideoExtractEventMessage struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	VideoUrl      string `json:"video_url"`
	IsUseSubTitle bool   `json:"is_use_subtitle"`
}

type VideoDetail struct {
	ChannelName string
	Title       string
	Description string
	Thumbnails  string
	PublishAt   string
	Duration    string
	ViewCount   uint64
	LikeCount   uint64
}

type ChannelDetail struct {
	ChannelName string
}
