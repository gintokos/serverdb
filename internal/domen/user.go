package domen

import "time"

type User struct {
	TelegramID int64 `gorm:"primaryKey"`
	CreatedAt  time.Time
}
