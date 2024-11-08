package domen

import "time"

type User struct {
	TelegramID uint `gorm:"primaryKey"`
	CreatedAt  time.Time
}