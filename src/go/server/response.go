package server

import (
  "encoding/json"
  "github.com/google/jsonapi"
  "log"
  "net/http"
)

// RawJsonResp is exported
func RawJsonResp(res http.ResponseWriter, statusCode int, data interface{}) {
  res.Header().Set("Content-Type", "application/json")
  res.WriteHeader(statusCode)
  err := json.NewEncoder(res).Encode(data)
  if err != nil {
    log.Print("JSONResp: error during encoding: ", err.Error())
    http.Error(res, "Error", http.StatusInternalServerError)
  }
}

// ModelJsonResp is exported
func ModelJsonResp(res http.ResponseWriter, req *http.Request, statusCode int, contextKey string) {
  model := req.Context().Value(contextKey)

  res.WriteHeader(statusCode)
  serializeErr := jsonapi.MarshalPayload(res, model)
  if serializeErr != nil {
    http.Error(res, "Error", http.StatusInternalServerError)
  }
}
