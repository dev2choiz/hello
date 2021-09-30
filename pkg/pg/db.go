package pg

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pg *gorm.DB

// NewDB return DB gorm instance
// should be called by wire only
func NewDB() *gorm.DB {
	if pg != nil {
		return pg
	}

	c := config.Conf
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Paris",
		c.PostgresHost,
		c.PostgresUser,
		c.PostgresPassword,
		c.PostgresDB,
		c.PostgresPort)
	pg, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf(err.Error())
	}
	logger.Info("bd initialized")
	return pg
}

func GetDB() *gorm.DB {
	return pg
}
