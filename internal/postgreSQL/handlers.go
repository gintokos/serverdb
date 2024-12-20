package postgresql

import (
	"errors"
	"fmt"
	"time"

	"github.com/gintokos/serverdb/internal/domain"

	"gorm.io/gorm"
)

func (postgre *PostgreSql) CreateUserRecord(id int64) error {
	db := postgre.db
	if id < 0 {
		return fmt.Errorf("userid must be greater than 0")
	}

	_, err := postgre.GetUserRecord(id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return fmt.Errorf("record with this id was created before")
	}

	record := domain.User{
		TelegramID: id,
		CreatedAt:  time.Now(),
	}
	if err := db.Create(&record).Error; err != nil {
		return err
	}

	postgre.logger.Info(fmt.Sprintf("Record with Telegram userID %d was sucsesfully created", id))
	return nil
}

func (postgre *PostgreSql) GetUserRecord(id int64) (domain.User, error) {
	db := postgre.db

	var user domain.User
	result := db.First(&user, id)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
