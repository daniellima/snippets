package app

import (
	"context"
	"net/http"
	"testing"

	"github.com/daniellima/counting-api/app/base/container"
	"github.com/daniellima/counting-api/app/base/test"
	"github.com/daniellima/counting-api/app/base/test/assert"
	"github.com/daniellima/counting-api/app/services"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	client      *test.Client
)

func TestMain(m *testing.M) {

	container.SetProvider("counterService", func(c *container.Container) interface{} {
		return services.NewRedisService(context.Background(), "redis://redis:6379")
	})
	container.SetProvider("logger", func(c *container.Container) interface{} {
		return &test.NopLogger{}
	})
	container.SetProvider("metricsService", func(c *container.Container) interface{} {
		return &test.NopMetricsService{}
	})

	redisClient = redis.NewClient(&redis.Options{Addr: "redis:6379"})

	serveMux := http.NewServeMux()
	ConfigureRoutes(serveMux)
	client = test.NewClient(serveMux)

	m.Run()
}

func setupTest(t *testing.T) {
	_, err := redisClient.FlushAll(context.Background()).Result()
	if err != nil {
		t.Errorf("Error when cleaning database for tests: %v", err)
	}
}

func assertCounter(t *testing.T, expectedCounter string) {
	actualCounter, _ := redisClient.Get(context.Background(), "counter").Result()

	if expectedCounter != actualCounter {
		t.Errorf("Counter \"%v\" is different from expected counter \"%v\"", actualCounter, expectedCounter)
	}
}

func Test_get_api_v1_Should_return_hello_world(t *testing.T) {
	setupTest(t)

	recorder := client.RequestGET("/api/v1")

	assert.Status(t, 200, recorder)
	assert.Body(t, `"Hello World"`, recorder)
}

func Test_get_readyz_Should_return_status_ok(t *testing.T) {
	setupTest(t)

	recorder := client.RequestGET("/readyz")

	assert.Status(t, 200, recorder)
	assert.Body(t, "", recorder)
}

func Test_get_livez_Should_return_status_ok(t *testing.T) {
	setupTest(t)

	recorder := client.RequestGET("/livez")

	assert.Status(t, 200, recorder)
	assert.Body(t, "", recorder)
}

func Test_get_api_v1_count_When_database_empty_Should_return_1(t *testing.T) {
	setupTest(t)

	recorder := client.RequestGET("/api/v1/count")

	assert.Status(t, 200, recorder)
	assert.Body(t, `{"count":1}`, recorder)
}

func Test_get_api_v1_count_When_called_Should_return_current_counter_plus_1(t *testing.T) {
	setupTest(t)

	redisClient.Set(context.Background(), "counter", 5, 0)

	recorder := client.RequestGET("/api/v1/count")

	assert.Status(t, 200, recorder)
	assert.Body(t, `{"count":6}`, recorder)
	assertCounter(t, "6")
}

func Test_post_api_v1_count_When_called_Should_set_counter(t *testing.T) {
	setupTest(t)

	redisClient.Set(context.Background(), "counter", 5, 0)

	recorder := client.RequestPOST("/api/v1/count", `{"count": 11}`)

	assert.Status(t, 200, recorder)
	assert.Body(t, "", recorder)
	assertCounter(t, "11")
}

func Test_post_api_v1_count_When_called_with_invalid_json_Should_return_bad_request(t *testing.T) {
	setupTest(t)

	redisClient.Set(context.Background(), "counter", 5, 0)

	recorder := client.RequestPOST("/api/v1/count", `{"count: 11}`)

	assert.Status(t, 400, recorder)
	assert.Body(t, `{"error":"unexpected EOF"}`, recorder)
	assertCounter(t, "5")
}

func Test_post_api_v1_count_When_called_with_counter_with_wrong_type_Should_return_bad_request(t *testing.T) {
	setupTest(t)

	redisClient.Set(context.Background(), "counter", 5, 0)

	recorder := client.RequestPOST("/api/v1/count", `{"count": "11"}`)

	assert.Status(t, 400, recorder)
	assert.Body(t, `{"error":"json: cannot unmarshal string into Go struct field CountHandlerRequestData.Count of type int"}`, recorder)
	assertCounter(t, "5")
}
