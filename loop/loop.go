package loop

import (
	"context"
	"errors"
	"log/slog"

	"github.com/lyricat/goutils/ai"
)

type (
	Message struct {
		ReplyContent string
		Context      context.Context
		BotID        uint64
		Lang         string

		UserIdentity string
		Content      string
		ConvKey      string

		DoneChan chan struct{}
	}

	Result struct {
		Response string
		Err      error
	}

	LoopService struct {
		aiInst  *ai.Instant
		adapter Adapter
	}

	Adapter interface {
		GetName() string
		GetMessageChan(ctx context.Context) <-chan *Message
		BeforeHandlingResult(req *Message)
		HandleResult(message *Message, result *Result)
	}
)

func New(aiInst *ai.Instant, adapter Adapter) *LoopService {
	return &LoopService{
		aiInst:  aiInst,
		adapter: adapter,
	}
}

func (s *LoopService) GetAdapterName() string {
	return s.adapter.GetName()
}

func (s *LoopService) Start(ctx context.Context) error {
	msgChan := s.adapter.GetMessageChan(ctx)
	for {
		select {
		case msg := <-msgChan:
			s.adapter.BeforeHandlingResult(msg)

			resp, err := s.HandleMessage(ctx, msg)
			if err != nil {
				slog.Error("[messageloop] failed to handle message", "error", err)
			} else {
				s.adapter.HandleResult(msg, &Result{
					Response: resp,
					Err:      err,
				})
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *LoopService) HandleMessage(ctx context.Context, m *Message) (string, error) {
	rp := ai.SusanoParams{
		Format: "json",
		Conditions: ai.SusanoParamsConditions{
			PreferredModel: "o1-mini",
		},
	}
	ret, err := s.aiInst.OneTimeRequestWithParams(ctx, m.Content, rp.ToMap())
	if err != nil {
		return "", err
	}
	if val, ok := ret.Json["response"]; ok {
		return val.(string), nil
	}
	return "", errors.New("response not found")
}
