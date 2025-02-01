package gen

import (
	"github.com/quailyquaily/gzk9000/store"

	_ "github.com/quailyquaily/gzk9000/store/agent"
	_ "github.com/quailyquaily/gzk9000/store/memslice"
	_ "github.com/quailyquaily/gzk9000/store/fact"
	_ "github.com/quailyquaily/gzk9000/store/studygoal"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen",
		Short: "generate database operation code",
		Run: func(cmd *cobra.Command, args []string) {

			h := store.WithContext(cmd.Context())
			h.Generate()
		},
	}

	return cmd
}
