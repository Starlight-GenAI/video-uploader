package repo

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/dreammnck/video-uploader/pkg/v1/model"
)

type queueHistoryFirestoreRepo struct {
	collection *firestore.CollectionRef
}

func NewQueueHistoryFirestoreRepo(client *firestore.Client, collectionName string) *queueHistoryFirestoreRepo {
	collection := client.Collection(collectionName)
	return &queueHistoryFirestoreRepo{collection: collection}
}

var _ model.IQueueHistoryFirestoreRepo = new(queueHistoryFirestoreRepo)

func (q *queueHistoryFirestoreRepo) Insert(ctx context.Context, data model.QueueHistory) error {

	if _, _, err := q.collection.Add(ctx, model.QueueHistoryFirestore{
		ID:            data.ID,
		UserID:        data.UserID,
		VideoUrl:      data.VideoUrl,
		VideoID:       data.VideoID,
		Title:         data.Title,
		Description:   data.Description,
		Thumbnails:    data.Thumbnails,
		Status:        data.Status,
		CreatedAt:     data.CreatedAt,
		UpdatedAt:     data.UpdatedAt,
		ChannelName:   data.ChannelName,
		IsUseSubTitle: data.IsUseSubTitle,
	}); err != nil {
		return err
	}

	return nil
}
