package migrations

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

var Collection = migrations.NewCollection()

func init() {
	Collection.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println(`Create data table`)
		_, err := db.Exec(`CREATE TABLE "data" ("id" bigserial,"field1" text,"field2" text,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,PRIMARY KEY ("id"))`)
		if err != nil {
			return err
		}

		fmt.Println(`CREATE INDEX "idx_data_deleted_at" ON "data.deleted_at"`)
		_, err = db.Exec(`CREATE INDEX "idx_data_deleted_at" ON "data" ("deleted_at")`)
		return err

	}, func(db migrations.DB) error {
		return nil
	})
}