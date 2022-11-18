package main

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table fundraising...")
		_, err := db.Exec(`CREATE TABLE fundraising(
      id SERIAL PRIMARY KEY,
      description TEXT NOT NULL,
      payment_info TEXT NOT NULL,
      social_link TEXT NOT NULL,
      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      closed BOOLEAN NOT NULL DEFAULT FALSE
    )`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table fundraising...")
		_, err := db.Exec(`DROP TABLE fundraising`)
		return err
	})
}
