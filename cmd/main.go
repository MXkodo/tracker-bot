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
			Text:  "Приложение",
			WebApp: webAppInfo,
		}

		inlineKeys := [][]tele.InlineButton{
			{btn},
		}
		inlineMarkup := &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		}

		if _, err := bot.Send(c.Chat(), "Здравствуйте! Нажмите на кнопку чтобы посмотреть список задач.", tele.ModeHTML, inlineMarkup); err != nil {
			return err
		}

		return nil
	})
	bot.Start()
}
