package services

import (
	"context"
	"fmt"

	"github.com/daniellima/counting-api/app/base/container"
	redis "github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
)

type RedisService struct {
	redisClient *redis.Client
}

func NewRedisService(ctx context.Context, redisConnectionString string) *RedisService {
	opt, err := redis.ParseURL(redisConnectionString)
	if err != nil {
		container.GetLogger().LogError(ctx, fmt.Sprintf("Could not parse redis connection string: %v", err), err)
	}

	return &RedisService{
		redisClient: redis.NewClient(opt),
	}
}

func (this *RedisService) SetCounter(ctx context.Context, newCounter int) error {
	ctx, span := otel.Tracer("onboarding-counter-api").Start(ctx, "RedisService.SetCounter")
	defer span.End()

	_, err := this.redisClient.Set(ctx, "counter", newCounter, 0).Result()

	return err
}

func (this *RedisService) IncrementCounter(ctx context.Context) (int, error) {
	ctx, span := otel.Tracer("onboarding-counter-api").Start(ctx, "RedisService.IncrementCounter")
	defer span.End()

	value, err := this.redisClient.Incr(ctx, "counter").Result()
	return int(value), err
}
