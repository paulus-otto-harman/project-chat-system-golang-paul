package domain

type Room struct {
	ID   uint    `gorm:"primaryKey,autoIncrement" json:"id"`
	Name *string `json:"name"`

	Users []User `gorm:"many2many:room_users;"`
}
