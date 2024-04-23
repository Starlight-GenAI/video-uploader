package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/video-uploader/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *videoUploaderHandler) VideoInfo(c echo.Context) error {
	body := new(serializer.VideoInfoRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	video_detail, err := h.videoUploaderSvc.GetVideoDetail(c.Request().Context(), body.VideoUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if video_detail == nil {
		return c.JSON(http.StatusNotFound, "not found video")
	}

	return c.JSON(http.StatusOK, serializer.VideoInfoResponse{
		VideoDetail: serializer.VideoDetail{
			Title:       video_detail.Title,
			ChannelName: video_detail.ChannelName,
			Description: video_detail.Description,
			Thumbnails:  video_detail.Thumbnails,
			PublishAt:   video_detail.PublishAt,
			Duration:    video_detail.Duration,
			ViewCount:   video_detail.ViewCount,
			LikeCount:   video_detail.LikeCount,
		},
	})

}
