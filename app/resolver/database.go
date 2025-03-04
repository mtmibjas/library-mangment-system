package resolver

import (
	"context"
	"database/sql"
	"fmt"
	"library-mngmt/app/config"
	"library-mngmt/pkg/logger/zap"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func resolveDBAdapter(cfg config.DBConfig) (*sql.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.PoolSize)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")

	return db, nil
}

func resolveRedisAdapter(cfg config.Redis) *redis.Client {

	db := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	result, err := db.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error pinging Redis server:", err)
		panic(err)
	}
	fmt.Println("Redis server ping result:", result)
	zap.Debug("Redis connect:", "Redis client connected successfully")

	return db
}
