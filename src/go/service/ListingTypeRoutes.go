package service

import (
  "fmt"
  "github.com/AdityaHegde/PathOfExileTrade/server"
  "github.com/google/jsonapi"
  "github.com/gorilla/mux"
  "gorm.io/gorm"
  "net/http"
)

// HandleGetAllListingType is exported
func HandleGetAllListingType(res http.ResponseWriter, db *gorm.DB) {
  listings, err := GetAllListingType(db)

  if err != nil {
    fmt.Println(err)
    http.Error(res, "Error", http.StatusInternalServerError)
  } else {
    server.ModelJsonResp(res, 200, listings)
  }
}

// HandleGetAllActiveListingType is exported
func HandleGetAllActiveListingType(res http.ResponseWriter, db *gorm.DB) {
  listings, err := GetAllActiveListingType(db)

  if err != nil {
    fmt.Println(err)
    http.Error(res, "Error", http.StatusInternalServerError)
  } else {
    server.ModelJsonResp(res, 200, listings)
  }
}

// HandleIDListingType is exported
func HandleIDListingType(res http.ResponseWriter, req *http.Request, db *gorm.DB) {
  var resp interface{}
  vars := mux.Vars(req)

  var reqListing = new(ListingType)
  reqParseErr := jsonapi.UnmarshalPayload(req.Body, reqListing)
  if (req.Method == http.MethodPost || req.Method == http.MethodPut) && reqParseErr != nil {
    fmt.Println(reqParseErr)
    http.Error(res, "Error", http.StatusInternalServerError)
    return
  }

  switch req.Method {
  case http.MethodGet:
    getListing, getErr := GetListingType(db, vars["id"])
    if getErr != nil {
      fmt.Println(getErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    resp = getListing

  case http.MethodPost:
    createListing, createErr := CreateListingType(db, reqListing)
    if createErr != nil {
      fmt.Println(createErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    resp = createListing

  case http.MethodPut:
    updateListing, updateErr := UpdateListingType(db, reqListing)
    if updateErr != nil {
      fmt.Println(updateErr)
      http.Error(res, "Error", http.StatusInternalServerError)
      return
    }
    resp = updateListing
  }

  server.ModelJsonResp(res, 200, resp)
}
