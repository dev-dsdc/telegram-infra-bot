package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/dev-dsdc/telegram-infra-bot/internal/health"
)

func main() {
	_ = godotenv.Load()

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("❌ BOT_TOKEN not set in environment or .env")
	}
	// Запускаем health check
	health.Start()

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot_1: %v", err)
	}

	log.Printf("✅ Bot authorized on account: %s", bot.Self.UserName)
	log.Println("🚀 Bot started successfully!")

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			var reply string

			cmd := strings.ToLower(update.Message.Command())
			switch cmd {

			case "start":
				reply = "👋 Привет! Я инфраструктурный бот.\nПока умею немного, но скоро буду помогать с автоматизацией!"
			case "help":
				reply = "📘 Доступные команды:\n/start — приветствие\n/help — список команд\n/status - статус серверов"
			case "status":
				reply = "📊 Всё работает нормально ✅"
			default:
				reply = "❓ Неизвестная команда. Напиши /help, чтобы узнать, что я умею."
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}

	}()
	<-ctx.Done()
	log.Println("🛑 Shutting down gracefully...")

	time.Sleep(2 * time.Second)
	log.Println("✅ Завершено корректно")
}
