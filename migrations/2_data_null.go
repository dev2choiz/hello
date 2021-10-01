package migrations

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {
	Collection.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println(`Migration 2`)
		return nil

	}, func(db migrations.DB) error {
		return nil
	})
}