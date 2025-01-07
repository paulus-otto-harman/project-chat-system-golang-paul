package domain

type Room struct {
	ID   uint    `gorm:"primaryKey,autoIncrement" json:"id" swaggerignore:"true"`
	Name *string `json:"name" example:"Alumni Lumoshive Academy"`

	Users []User `gorm:"many2many:room_users;"`
}
