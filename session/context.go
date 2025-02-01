package session

import (
	"context"
)

type (
	contextKey struct{}
)

type (
	Session struct {
	}
)

func With(ctx context.Context, s *Session) context.Context {
	return context.WithValue(ctx, contextKey{}, s)
}

func From(ctx context.Context) *Session {
	return ctx.Value(contextKey{}).(*Session)
}
