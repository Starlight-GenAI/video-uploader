package handler

import "github.com/dreammnck/video-uploader/pkg/v1/model"

type videoUploaderHandler struct {
	videoUploaderSvc model.IVideoUplaoderSvc
}

func NewVideoUploaderHandler(videoUploaderSvc model.IVideoUplaoderSvc) *videoUploaderHandler {
	return &videoUploaderHandler{videoUploaderSvc}
}
