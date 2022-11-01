package respond

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daniellima/counting-api/app/base/container"
)

func logError(ctx context.Context, err error) {
	container.GetLogger().LogError(ctx, fmt.Sprintf("An error has ocurred when handling a request: %v", err), err)
}

func InternalError(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	logError(ctx, err)

	w.Write([]byte("\"A problem occurred. We are looking into it\""))
}

func BadRequestError(ctx context.Context, w http.ResponseWriter, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	logError(ctx, err)

	w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
}

func JSON(ctx context.Context, w http.ResponseWriter, v any) {
	w.Header().Set("Content-type", "application/json")

	response, err := json.Marshal(v)
	if err != nil {
		InternalError(ctx, w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
