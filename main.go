package main

import (
	"context"
	"fmt"
	"os"

	"github.com/quailyquaily/gzk9000/cmd"
	"github.com/quailyquaily/gzk9000/session"
)

func main() {
	ctx := context.Background()
	s := &session.Session{}
	ctx = session.With(ctx, s)
	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
