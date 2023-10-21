package mariadb

import (
	"database/sql"
	"fmt"
	databases "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/core"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDatabase struct {
	client *gorm.DB
	logger zap.Logger
}

func NewMariaDBInstance(log, config interface{}) (*MariaDatabase, error) {
	if config == nil {
		return &MariaDatabase{}, fmt.Errorf("MariaDB config cannot be nil")
	}
	cnf := config.(databases.MariaDBConfig)
	mysqlClient, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", cnf.Username, cnf.Password, cnf.Host, cnf.Database))
	if err != nil {
		return &MariaDatabase{}, fmt.Errorf("error initialising sql connection: %v", err)
	}
	mDBGormClient, err := gorm.Open(mysql.New(mysql.Config{Conn: mysqlClient}), &gorm.Config{})
	if err != nil {
		return &MariaDatabase{}, fmt.Errorf("error initialising gorm client with mysql client: %v", err)
	}
	sqlDB, err := mDBGormClient.DB()
	if err != nil {
		return &MariaDatabase{}, fmt.Errorf("error initialising maria db client: %v", err)
	}
	sqlDB.SetMaxIdleConns(cnf.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cnf.MaxOpenConnections)
	return &MariaDatabase{
		client: mDBGormClient,
		logger: log.(zap.Logger),
	}, nil
}

func (rd *MariaDatabase) Create(config, data interface{}) error             { return nil }
func (rd *MariaDatabase) Get(config, data interface{}) (interface{}, error) { return nil, nil }
func (rd *MariaDatabase) Update(config, data interface{}) error             { return nil }
func (rd *MariaDatabase) Delete(config, data interface{}) error             { return nil }
