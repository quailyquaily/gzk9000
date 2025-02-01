package migrate

import (
	"fmt"
	"strconv"

	"github.com/lyricat/goutils/qdrant"
	"github.com/quailyquaily/gzk9000/store"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database tables",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "Migrate the DB to the most recent version available",
		RunE: func(cmd *cobra.Command, args []string) error {
			return store.WithContext(cmd.Context()).MigrationUp()
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "up-to VERSION",
		Short: "Migrate the DB to a specific VERSION",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("up-to requires a version argument")
			}
			version, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid version: %w", err)
			}
			return store.WithContext(cmd.Context()).MigrationUpTo(version)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "down",
		Short: "Roll back the version by 1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return store.WithContext(cmd.Context()).MigrationUp()
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "down-to VERSION",
		Short: "Roll back to a specific VERSION",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("up-to requires a version argument")
			}
			version, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid version: %w", err)
			}
			return store.WithContext(cmd.Context()).MigrationDownTo(version)
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "redo",
		Short: "Re-run the latest migration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return store.WithContext(cmd.Context()).MigrationRedo()
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Dump the migration status for the current DB",
		RunE: func(cmd *cobra.Command, args []string) error {
			return store.WithContext(cmd.Context()).MigrationStatus()
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "create NAME",
		Short: "Creates new migration file with the current timestamp",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("create requires a name argument")
			}
			return store.WithContext(cmd.Context()).MigrationCreate(args[0])
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use:   "init-vector",
		Short: "Initialize vector tables",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			qd := qdrant.New(qdrant.Config{
				Addr:   viper.GetString("qdrant.addr"),
				APIKey: viper.GetString("qdrant.api_key"),
			})
			if err := qd.CreateCollection(ctx, qdrant.CreateCollectionParams{
				CollectionName: viper.GetString("sys.collections.fact"),
				VectorSize:     1536,
				Indexes: []qdrant.CreateCollectionIndexItem{{
					Name: "fact_id", Type: "int",
				}, {
					Name: "agent_id", Type: "int",
				}},
			}); err != nil {
				return err
			}
			return nil
		},
	})

	return cmd
}
