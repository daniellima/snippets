package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daniellima/counting-api/app/base/container"
	"github.com/daniellima/counting-api/app/base/respond"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	counterService := container.GetCounterService()

	_, err := counterService.IncrementCounter(r.Context())
	if err != nil {
		respond.InternalError(r.Context(), w, err)
	} else {
		respond.JSON(r.Context(), w, "Hello World")
	}
}

func ReadyzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type CountHandlerRequestData struct {
	Count int
}

func ShowCountHandler(w http.ResponseWriter, r *http.Request) {
	counterService := container.GetCounterService()

	value, err := counterService.IncrementCounter(r.Context())
	if err != nil {
		respond.InternalError(r.Context(), w, err)
		return
	}

	respond.JSON(r.Context(), w, map[string]int{
		"count": value,
	})
}

func UpdateCountHandler(w http.ResponseWriter, r *http.Request) {
	counterService := container.GetCounterService()

	dec := json.NewDecoder(r.Body)
	requestData := CountHandlerRequestData{}
	err := dec.Decode(&requestData)
	if err != nil {
		respond.BadRequestError(r.Context(), w, err)
		return
	}

	container.GetLogger().Log(r.Context(), fmt.Sprintf("New count value received: %v", requestData.Count))

	err = counterService.SetCounter(r.Context(), requestData.Count)
	if err != nil {
		respond.InternalError(r.Context(), w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
