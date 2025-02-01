package migrate

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/quailyquaily/gzk9000/store"

	"github.com/spf13/cobra"
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
	// temp
	cmd.AddCommand(&cobra.Command{
		Use:   "fill-payee-id",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			tx := store.WithContext(ctx)

			lps := make([]struct {
				ID         uint64
				ListID     uint64
				EvmAddress string
			}, 0)
			if err := tx.Raw("select id, list_id, evm_address from list_payments;").Scan(&lps).Error; err != nil {
				slog.Error("[migrate] failed to get list and evm_address", "error", err)
				return err
			}

			var result struct {
				ID     uint64
				UserID uint64
			}
			for _, lp := range lps {
				if err := tx.Raw(fmt.Sprintf("select id, user_id from lists where id =%d limit 1;", lp.ListID)).Scan(&result).Error; err != nil {
					slog.Error("[migrate] failed to get list id and user id", "error", err)
					return err
				}
				if result.ID != 0 {
					// tx.Exec(fmt.Sprintf("update list_payments set payee_id = %d where id = %d;", result.UserID, lp.ID))
					sql := fmt.Sprintf("update user_payouts set evm_address ='%s' where user_id = %d;", lp.EvmAddress, result.UserID)
					slog.Info("[migrate] update sql", "sql", sql)
					tx.Exec(sql)
				}
			}
			return nil
		},
	})

	return cmd
}
