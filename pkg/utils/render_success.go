package utils

import (
	"encoding/json"
	"github.com/renanmachad/ecommerce-rest/pkg/app/dto/response"
	"net/http"
)

func RenderSuccess(w http.ResponseWriter, r *http.Request, successResponse *response.SuccessResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsonBytes, err := json.Marshal(successResponse)
	ErrorHandler(err)

	_, err = w.Write(jsonBytes)
	ErrorHandler(err)
	return
}
