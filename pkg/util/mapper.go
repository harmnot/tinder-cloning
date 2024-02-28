package util

import (
	"encoding/json"
	"net/http"
)

type MapResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	d, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = w.Write(d)
	return err
}

func RenderJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	if val, isErr := v.(error); isErr {
		_ = WriteJSON(w, statusCode, val)
		return
	}
	resp := MapResponse{
		Message: http.StatusText(statusCode),
		Data:    v,
	}
	_ = WriteJSON(w, statusCode, resp)
}
