package serializer

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UploadVideoRequest struct {
	VideoUrl      string `json:"video_url"`
	VideoID       string `json:"video_id"`
	IsUseSubtitle *bool  `json:"is_use_subtitle"`
	UserID        string `json:"user_id"`
}

type UploadVideoResponse struct {
	ID string `json:"queue_id"`
}

func (b UploadVideoRequest) Validate() error {
	if err := v.ValidateStruct(&b,
		v.Field(&b.IsUseSubtitle, v.NotNil),
		v.Field(&b.VideoUrl, v.When(b.VideoID == "", v.Required, is.URL)),
		v.Field(&b.VideoID, v.When(b.VideoUrl == "", v.Required)),
		v.Field(&b.UserID, v.Required),
	); err != nil {
		return err
	}
	return nil
}
