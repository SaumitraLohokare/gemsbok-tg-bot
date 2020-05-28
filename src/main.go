package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL") // you must add it to your config vars
		token     = os.Getenv("TOKEN")      // you must add it to your config vars
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, "Type \n/add : to add to your account\n/sub : to subtract from your account")
	})

	b.Handle("/add", func(m *tb.Message) {
		if m.Sender.Username != "saumi_l" {
			b.Send(m.Sender, "Sorry but you are not registered for the service.")
			return
		}

		msg := ""
		command_args := strings.Split(m.Text, " ")
		if len(command_args) < 3 {
			msg = "Usage for /add command is:\n/add <amount> <comment>"
			b.Send(m.Sender, msg)
			return
		}

		amt, err := strconv.ParseFloat(command_args[1], 64)
		if err != nil {
			msg = "Usage for /add command is:\n/add <amount> <comment>"
			b.Send(m.Sender, msg)
			return
		}

		comment := strings.Join(command_args[2:], " ")

		timestamp := time.Now()

		msg = "Added to your account: " + fmt.Sprintf("%f", amt) + "\nwith comment: " + comment + "\nTime: " + timestamp.Format("2006-01-02 15:04:05")

		b.Send(m.Sender, msg)
	})

	b.Handle("/sub", func(m *tb.Message) {
		if m.Sender.Username != "saumi_l" {
			b.Send(m.Sender, "Sorry but you are not registered for the service.")
			return
		}

		msg := ""
		command_args := strings.Split(m.Text, " ")
		if len(command_args) < 3 {
			msg = "Usage for /sub command is:\n/sub <amount> <comment>"
			b.Send(m.Sender, msg)
			return
		}

		amt, err := strconv.Atoi(command_args[1])
		if err != nil {
			msg = "Usage for /sub command is:\n/sub <amount> <comment>"
			b.Send(m.Sender, msg)
			return
		}

		comment := strings.Join(command_args[2:], " ")

		timestamp := time.Now()

		msg = "Subtracted from your account: " + strconv.Itoa(amt) + "\nwith comment: " + comment + "\nTime: " + timestamp.Format("2006-01-02 15:04:05")
		b.Send(m.Sender, msg)
	})

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Type /help for commands")
	})

	b.Start()
}
