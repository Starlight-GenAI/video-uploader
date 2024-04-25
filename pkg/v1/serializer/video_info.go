package serializer

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type VideoDetail struct {
	Title       string `json:"title"`
	ChannelName string `json:"channel_name"`
	Description string `json:"decription"`
	Thumbnails  string `json:"thumbnails"`
	PublishAt   string `json:"publish_at"`
	Duration    string `json:"duration"`
	ViewCount   uint64 `json:"view_count"`
	LikeCount   uint64 `json:"like_count"`
}

type VideoInfoRequest struct {
	VideoID string `json:"video_id"`
}

type VideoInfoResponse struct {
	VideoDetail VideoDetail `json:"video_detail"`
}

func (b VideoInfoRequest) Validate() error {
	if err := v.ValidateStruct(&b,
		v.Field(&b.VideoID, v.Required),
	); err != nil {
		return err
	}
	return nil
}
