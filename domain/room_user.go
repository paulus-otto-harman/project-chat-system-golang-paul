package domain

type RoomUser struct {
	ID     uint `gorm:"primaryKey,autoIncrement"`
	RoomID uint
	UserID uint
}
