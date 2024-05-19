package model

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const uri_REDIS = "redis:6379"

type REDIS struct {
	client *redis.Client
}

func NewREdis(ctx context.Context) *REDIS {
	rdb := redis.NewClient(&redis.Options{
		Addr:     uri_REDIS,
		Password: "",
		DB:       0,
	})
	return &REDIS{client: rdb}
}
