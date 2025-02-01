package telegram

import (
	"context"
	"log/slog"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/quailyquaily/gzk9000/core"
	"github.com/quailyquaily/gzk9000/loop"
)

type (
	messageKey struct{}
)

type (
	Config struct {
		Token string
		Debug bool
	}
	service struct {
		name   string
		cfg    Config
		client *tgbotapi.BotAPI
	}
)

func New(
	cfg Config,
	memslices core.MemsliceStore,
	factz core.FactService,
) *service {
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		return nil
	}
	bot.Debug = cfg.Debug

	return &service{
		name:   "telegram",
		cfg:    cfg,
		client: bot,
	}
}

func (s *service) GetName() string {
	return s.name
}

func (s *service) GetMessageChan(ctx context.Context) <-chan *loop.Message {
	msgChan := make(chan *loop.Message)
	go func() {
		u := tgbotapi.NewUpdate(0)
		updates := s.client.GetUpdatesChan(u)
		for update := range updates {
			if update.Message == nil || update.Message.Chat == nil || update.Message.Text == "" {
				continue
			}

			prefix := "@" + s.client.Self.UserName
			if update.Message.Chat.IsGroup() || update.Message.Chat.IsSuperGroup() {
				if update.Message.ReplyToMessage == nil || update.Message.ReplyToMessage.From.ID != s.client.Self.ID {
					if !strings.HasPrefix(update.Message.Text, prefix) {
						continue
					}
				}
			}
			replyContent := ""
			if update.Message.ReplyToMessage != nil {
				replyContent = update.Message.ReplyToMessage.Text
			}

			content := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, prefix))
			messageCtx := context.WithValue(ctx, messageKey{}, update.Message)
			msgChan <- &loop.Message{
				Context:      messageCtx,
				ReplyContent: replyContent,
				Content:      content,
				UserIdentity: strconv.FormatInt(update.Message.From.ID, 10),
				ConvKey:      strconv.FormatInt(update.Message.Chat.ID, 10),
			}
		}
		<-ctx.Done()
		s.client.StopReceivingUpdates()
		close(msgChan)
	}()

	return msgChan
}

func (b *service) BeforeHandlingResult(req *loop.Message) {
	msg := req.Context.Value(messageKey{}).(*tgbotapi.Message)
	// set the bot as "typing" before sending the message
	typingAction := tgbotapi.NewChatAction(msg.Chat.ID, tgbotapi.ChatTyping)
	if _, err := b.client.Send(typingAction); err != nil {
		slog.Warn("[telegram] failed to send typing action", "error", err)
	}
}

func (b *service) HandleResult(req *loop.Message, r *loop.Result) {
	if r.Err != nil {
		slog.Warn("[telegram] ignore error", "error", r.Err)
		return
	}
	text := ""
	if r.Err != nil {
		text = r.Err.Error()
	} else {
		text = r.Response
	}

	msg := req.Context.Value(messageKey{}).(*tgbotapi.Message)

	// if the response is empty, do not send anything
	if text == "" {
		slog.Warn("[telegram] empty response")
		b.sendMessage(msg.Chat.ID, "Sorry, I lost my words, please say again later.")
		return
	}

	// if the response is longer than 4096 characters, split it into multiple messages
	// each split message should be a chunk less than 4096 characters
	// the split should be done by line
	if len(text) > 4000 {
		parts := loop.SplitMarkdown(text, 4000)
		for _, part := range parts {
			b.sendMessage(msg.Chat.ID, part)
		}
		// lines := strings.Split(text, "\n")
		// chunk := ""
		// for _, line := range lines {
		// 	if len(chunk)+len(line) > 4096 {
		// 		b.sendMessage(msg.Chat.ID, chunk)
		// 		chunk = ""
		// 	}
		// 	chunk += line + "\n"
		// }
		// if chunk != "" {
		// 	b.sendMessage(msg.Chat.ID, chunk)
		// }
	} else {
		b.sendMessage(msg.Chat.ID, text)
	}
}

func (b *service) sendMessage(chatID int64, text string) {
	// send the message to the chat
	// disable web page preview
	slog.Info("[telegram] send message", "chat_id", chatID, "size", len(text))
	message := tgbotapi.NewMessage(chatID, text)
	message.ParseMode = tgbotapi.ModeMarkdown
	if _, err := b.client.Send(message); err != nil {
		slog.Error("[telegram] failed to send message", "error", err)
	}
}
