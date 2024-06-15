package app

import (
	"database/sql"
	"fmt"
	"foodV5/common/config"
	"foodV5/common/entity"
	mysqlDrive "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func NewGorm() (*gorm.DB, error) {
	var directory gorm.Dialector
	cfg, err := mysqlDrive.ParseDSN(config.C.MySQL.DSN())
	if err != nil {
		return nil, err
	}
	err = createDatabaseWithMySQL(cfg)
	if err != nil {
		return nil, err
	}
	directory = mysql.Open(config.C.MySQL.DSN())
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix:   config.C.Gorm.TablePrefix,
			SingularTable: true,
		},
	}
	db, err := gorm.Open(directory, gormConfig)
	if err != nil {
		return nil, err
	}

	if config.C.Gorm.Debug {
		db = db.Debug()
	}
	// 创建数据表
	err = db.AutoMigrate(
		&entity.Invite{},
		&entity.Integral{},
		&entity.User{},
		&entity.User{},
		&entity.Config{},
		&entity.UserAct{},
		&entity.UserFinance{},
	)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.C.Gorm.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.C.Gorm.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.C.Gorm.MaxLifetime) * time.Second)

	return db, nil
}

func createDatabaseWithMySQL(cfg *mysqlDrive.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)
	return err
}
