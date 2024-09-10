package main

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token: "6817719130:AAEeGjoL9bJg_E3LLn80VHwshVQDhMbgaWY",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

bot.Handle("/start", func(c tele.Context) error {
		chatID := c.Chat().ID  

		registerURL := fmt.Sprintf("https://taskfront.emivn.io?&chat_id=%d",chatID)

		webAppInfo := &tele.WebApp{
			URL: registerURL,
		}

		btn := tele.InlineButton{
			Text:  "📋",
			WebApp: webAppInfo,
		}

		inlineKeys := [][]tele.InlineButton{
			{btn},
		}
		inlineMarkup := &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		}

		if _, err := bot.Send(c.Chat(), "Здравствуйте! Добро пожаловать в EMIVN TRACKER.\nЭто приложение предназначено для управления вашими задачами.\nНажмите на кнопку ниже, чтобы просмотреть список задач и воспользоваться всеми функциями приложения.✅", tele.ModeHTML, inlineMarkup); err != nil {
			return err
		}

		return nil
	})
	bot.Start()
}
