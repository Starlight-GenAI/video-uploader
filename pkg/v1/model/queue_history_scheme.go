package model

import (
	"time"
)

type QueueHistoryFirestore struct {
	ID            string    `firestore:"queue_id"`
	UserID        string    `firestore:"user_id"`
	VideoUrl      string    `firestore:"video_url"`
	VideoID       string    `firestore:"video_id"`
	Status        string    `firestore:"status"`
	Title         string    `firestore:"title"`
	Description   string    `firestore:"description"`
	Thumbnails    string    `firestore:"thumbnails"`
	ChannelName   string    `firestore:"channel_name"`
	CreatedAt     time.Time `firestore:"created_at"`
	UpdatedAt     time.Time `firestore:"updated_at"`
	IsUseSubTitle bool      `firestore:"is_use_subtitle"`
}

type QueueHistory struct {
	ID            string
	UserID        string
	VideoUrl      string
	VideoID       string
	Status        string
	Title         string
	Description   string
	Thumbnails    string
	ChannelName   string
	IsUseSubTitle bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
