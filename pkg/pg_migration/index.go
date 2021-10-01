package pg_migration

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"log"
	"os"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

// Migrate execute migrations
func Migrate(params []string, conf *config.Config) error {
	if len(params) == 0 {
		logger.Info(usageText)
		return nil
	}

	// command check
	if err := check(params[0]); err != nil {
		return err
	}

	opt := &pg.Options{
		User:      conf.PostgresUser,
		Database:  conf.PostgresDB,
		Password:  conf.PostgresPassword,
		TLSConfig: nil,
	}
	var addr string
	log.Println(conf.AppEnvContext, os.Getenv("DB_SOCKET_DIR"))
	if conf.AppEnvContext == "cloud_function" {
		//opt.Addr = fmt.Sprintf("%s", conf.PostgresHost)
		//opt.Addr = fmt.Sprintf("%s:%s", conf.PostgresHost, conf.PostgresPort)
		opt.Addr = fmt.Sprintf("%s/.s.PGSQL.%s", conf.PostgresHost, conf.PostgresPort)
		opt.Network = "unix"
	} else {
		opt.Addr = fmt.Sprintf("%s:%s", conf.PostgresHost, conf.PostgresPort)
	}

	log.Println(addr)

	log.Println("in migration: before Connect()")
	db := pg.Connect(opt)
	defer db.Close()
	log.Println("in migration: after Connect()")

	oldVer, newVer, err := migrations.Run(db, params...)
	logResult(params[0], err, oldVer, newVer)
	return nil
}

func check(cmd string) error {
	switch cmd {
	case "init", "down", "up":
		return nil
	}
	return fmt.Errorf("'%s' command is unsupported", cmd)
}

func logResult(cmd string, err error, oldVer, newVer int64) {
	switch cmd {
	case "init":
		if err != nil {
			logger.Warnf("migration init failed: %s", err.Error())
			return
		}
		logger.Info("'gopg_migrations' created")
		return
	case "down":
	case "up":
		if err != nil {
			logger.Infof("migration %s failed: %s", cmd, err.Error())
			return
		}
		if newVer != oldVer {
			logger.Infof("migrated from version %d to %d", oldVer, newVer)
		} else {
			logger.Infof("nothing to play. current migration: %d", oldVer)
		}
		return
	default:
		logger.Infof("command %s is unsupported", cmd)
		return
	}
}
