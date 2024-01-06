package mariadb

import (
	"database/sql"
	"fmt"

	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MariaDatabase struct {
	client *gorm.DB
	logger logger.Logger
}

func NewMariaDBInstance(log logger.Logger, config interface{}) (*MariaDatabase, error) {
	if config == nil {
		return &MariaDatabase{}, fmt.Errorf("MariaDB config cannot be nil")
	}
	cnf := config.(models.DBConfig)

	mysqlClient, err := sql.Open("mysql", createConnectionString(cnf.Username, cnf.Password, cnf.Host, cnf.Database[0]))
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

	if err := mDBGormClient.AutoMigrate(&models.MariaDBPayload{}); err != nil {
		log.Errorf("Error creating user table in MariaDB")
	}

	return &MariaDatabase{
		client: mDBGormClient,
		logger: log,
	}, nil
}

func createConnectionString(user, password, host string, database interface{}) (connStr string) {

	if len(password) == 0 {
		connStr = fmt.Sprintf("%s@tcp(%s)/%s", user, host, database)
	} else {
		connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", user, password, host, database)
	}

	return
}

func (rd *MariaDatabase) Exists(data interface{}) bool {
	if data == nil {
		return false
	}
	// TODO
	return true
}

func (rd *MariaDatabase) Create(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("payload cannot be nil")
	}
	payload := data.(models.MariaDBPayload)

	return rd.client.Create(models.MariaDBPayload{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Username: payload.Username,
	}), nil
}

func (rd *MariaDatabase) Get(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("payload cannot be nil")
	}
	var user models.MariaDBPayload
	payload := data.(models.Fields)
	query := rd.client.Model(&models.MariaDBPayload{})
	for key, value := range payload {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	query.First(&user)
	return user, nil
}

func (rd *MariaDatabase) Update(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("payload cannot be nil")
	}
	var user models.MariaDBPayload
	queries := data.([]models.Fields)
	queryField, updateField := queries[0], queries[1]
	query := rd.client.Model(&models.MariaDBPayload{})
	for key, value := range queryField {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	query.Updates(updateField)
	query.First(&user)
	return user, nil
}

func (rd *MariaDatabase) Delete(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("payload cannot be nil")
	}
	payload := data.(models.Fields)
	query := rd.client.Model(&models.MariaDBPayload{})
	for key, value := range payload {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	if err := query.Delete(gorm.DB{}); err != nil {
		return nil, fmt.Errorf("failed to delete row %+v with: %s", payload, err.Error)
	}
	return nil, nil
}
