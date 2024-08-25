package postgres

import (
	"context"
	"fmt"
	"log"
	"time"
	configs "github.com/qara-qurt/telegram_plus/user_service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DatabasePSQL struct {
	DB *pgxpool.Pool
}

func Config(conf *configs.Database) *pgxpool.Config {
	const defaultMaxConn = int32(4)
	const defaultMinConn = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	URL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
		conf.SSLMode,
	)

	dbConfig, err := pgxpool.ParseConfig(URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConn
	dbConfig.MinConns = defaultMinConn
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return dbConfig
}

func NewDatabasePSQL(conf *configs.Database) (*DatabasePSQL, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config(conf))
	if err != nil {
		logrus.Errorf("error while creating connection to the database, %v", err)
		return nil, err
	}
	// conn, err := connPool.Acquire(context.Background())
	// if err != nil {
	// 	logrus.Errorf("error while accuring connection from the database pool, %v", err)
	// 	return nil, err
	// }
	// defer conn.Release()

	// PING database
	err = connPool.Ping(context.Background())
	if err != nil {
		logrus.Error("Could not ping database")
		return nil, err
	}

	logrus.Info("Connected to the database!!")

	return &DatabasePSQL{
		DB: connPool,
	}, nil
}
