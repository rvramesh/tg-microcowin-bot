package dao

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db "github.com/rvramesh/tg-microwin-bot/src/db"
	models "github.com/rvramesh/tg-microwin-bot/src/models"
)

var err error

func CheckAndGetUser(chat *tgbotapi.Chat) models.User {
	var user models.User
	var dbCon = db.GetDB()
	result := dbCon.Where("user_name = ?", chat.UserName).First(&user)
	if result.RowsAffected == 1 {
		return user
	} else {
		user := models.User{
			FirstName:    chat.FirstName,
			LastName:     chat.LastName,
			ChatId:       chat.ID,
			UserName:     chat.UserName,
			IsBlocked:    false,
			IsAdmin:      false,
			CurrentState: "Listen",
			Data:         "",
		}
		dbCon.Create(&user)
		fmt.Println(user.ID)
		return user
	}
}
