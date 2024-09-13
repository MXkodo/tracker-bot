package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tele "gopkg.in/telebot.v3"
)

const userAPIURL = "http://localhost:8081/api/v1/user" 

type UserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
	UUID     string `json:"uuid"`
	ChatID   int64  `json:"chat_id"`
}

func main() {
	pref := tele.Settings{
		Token:  "6817719130:AAEeGjoL9bJg_E3LLn80VHwshVQDhMbgaWY",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(c tele.Context) error {
		userName := c.Sender().Username
		if userName == "" {
			userName = c.Sender().FirstName 
		}

		user, err := getUserByUsername(userName)
		if err != nil {
			return err
		}

		var message string
		if user.Role == 0 {
			message = "Привет! Я EMIVN TRACKER — твой помощник в выполнении задач. Посмотрим, какие задачи ждут тебя сегодня? Нажми кнопку внизу, чтобы открыть приложение и начать работу."
		} else {
			message = "Привет! Я EMIVN TRACKER — ваш помощник по управлению задачами. Готов организовать ваши проекты. Какие задачи сегодня на повестке дня? Нажми кнопку внизу, чтобы открыть приложение и начать работу."
		}

		registerURL := fmt.Sprintf("https://taskfront.emivn.io?&chat_id=%d", c.Chat().ID)

		webAppInfo := &tele.WebApp{
			URL: registerURL,
		}

		btn := tele.InlineButton{
			Text:  "Начать",
			WebApp: webAppInfo,
		}

		inlineKeys := [][]tele.InlineButton{
			{btn},
		}
		inlineMarkup := &tele.ReplyMarkup{
			InlineKeyboard: inlineKeys,
		}

		if _, err := bot.Send(c.Chat(), message, tele.ModeHTML, inlineMarkup); err != nil {
			return err
		}

		return nil
	})

	bot.Start()
}

func getUserByUsername(username string) (*UserResponse, error) {
	requestBody, err := json.Marshal(map[string]string{"username": username})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(userAPIURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userResponse UserResponse
	if err := json.Unmarshal(body, &userResponse); err != nil {
		return nil, err
	}

	return &userResponse, nil
}
