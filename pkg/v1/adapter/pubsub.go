package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/dreammnck/video-uploader/pkg/v1/model"
)

type pubsubAdapter struct {
	topic *pubsub.Topic
}

func NewPubSubAdapter(client *pubsub.Client, topicName string) *pubsubAdapter {
	topic := client.Topic(topicName)
	return &pubsubAdapter{topic: topic}
}

var _ model.IPubSubAdapter = new(pubsubAdapter)

func (p *pubsubAdapter) Publish(ctx context.Context, data model.VideoExtractEventMessage) error {
	publishErr := make(chan error)

	dataBytes, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	result := p.topic.Publish(ctx, &pubsub.Message{
		Data: dataBytes,
	})

	go func(res *pubsub.PublishResult) {
		_, err := res.Get(ctx)
		if err != nil {
			publishErr <- errors.New(fmt.Sprintf("publish failed with %s", err.Error()))
		} else {
			publishErr <- nil
		}
	}(result)

	return <-publishErr
}
