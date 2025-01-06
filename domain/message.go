package domain

import "time"

type Message struct {
	ID            uint      `json:"id,-"`
	RoomID        uint      `json:"roomId,-"`
	SenderID      uint      `json:"sender,omitempty"`
	Content       string    `json:"content,omitempty"`
	AttachmentUrl string    `json:"attachmentUrl,omitempty"`
	ReplyTo       int       `json:"replyTo,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
