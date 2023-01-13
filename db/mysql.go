/*
 * @PackageName: db
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 10:36
 */

package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

type MysqlOptions struct {
	Source                string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	Log                   *log.Helper
	GormCfg               *gorm.Config
}

func NewMysql(opts *MysqlOptions) (*gorm.DB, error) {
	if opts.GormCfg == nil {
		opts.GormCfg = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true, //在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		}
	}
	if opts.Log == nil {
		opts.Log = log.NewHelper(log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"module", "mysql"))
	}
	db, err := gorm.Open(mysql.Open(opts.Source), opts.GormCfg)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	db.Use(&TracePlugin{
		log: opts.Log,
	})
	return db, nil
}
