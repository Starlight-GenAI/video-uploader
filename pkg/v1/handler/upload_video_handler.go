package handler

import (
	"fmt"
	"net/http"

	"github.com/dreammnck/video-uploader/pkg/v1/model"
	"github.com/dreammnck/video-uploader/pkg/v1/serializer"
	"github.com/labstack/echo/v4"
)

func (h *videoUploaderHandler) UploadVideoHandler(c echo.Context) error {

	body := new(serializer.UploadVideoRequest)
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot bind request body")
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("validate request fail with %s", err.Error()))
	}

	id, err := h.videoUploaderSvc.Publish(c.Request().Context(), model.PublishOpts{
		UserID:        body.UserID,
		IsUseSubTitle: *body.IsUseSubtitle,
		VideoUrl:      body.VideoUrl,
		VideoID:       body.VideoID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, serializer.UploadVideoResponse{ID: *id})
}
