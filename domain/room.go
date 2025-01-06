package domain

type Room struct {
	ID   uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Name *string `json:"name"`
}
