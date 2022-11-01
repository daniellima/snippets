package container

import (
	"log"
)

var c = New()

type Container struct {
	providerMap map[string]ProviderFunc
	serviceMap  map[string]interface{}
}

type ProviderFunc func(*Container) interface{}

func New() *Container {
	msp := make(map[string]ProviderFunc)
	msi := make(map[string]interface{})
	return &Container{msp, msi}
}

func SetProvider(name string, provider ProviderFunc) {
	c.providerMap[name] = provider
}

func Get(name string) interface{} {
	if service, ok := c.serviceMap[name]; ok {
		return service
	}

	if provider, ok := c.providerMap[name]; ok {
		service := provider(c)
		c.serviceMap[name] = service
		return service
	}

	log.Fatalf("Dependency %s not registered!", name)

	return nil
}

func GetConfig() Config {
	return Get("config").(Config)
}

func GetLogger() Logger {
	return Get("logger").(Logger)
}

func GetCounterService() CounterService {
	return Get("counterService").(CounterService)
}

func GetMetricsService() MetricsService {
	return Get("metricsService").(MetricsService)
}
