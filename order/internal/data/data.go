package data

import (
	"order/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewOrderRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	log := log.NewHelper(log.With(logger, "module", "order-service/data"))
	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}

/**
 * 创建DB
 */
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "order-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&Order{}); err != nil {
		log.Fatal(err)
	}
	return db
}
