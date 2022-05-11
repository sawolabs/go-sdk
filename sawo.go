package sawo

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var ApiKey string
var IdentifierType string
var FilePath string
var Route string



type SawoConfig struct{
	ApiKey string
	IdentifierType string
	FilePath string
	Route string
}

type SawoPayload struct {
	UserID                 string    `json:"user_id"`
	CreatedOn              time.Time `json:"created_on"`
	Identifier             string    `json:"identifier"`
	IdentifierType         string    `json:"identifier_type"`
	VerificationToken      string    `json:"verification_token"`
	CustomFieldInputValues struct {
	} `json:"customFieldInputValues"`
}

func (e *SawoConfig) Init(apikey string, identifiertype string, filePath string, route string) {
	ApiKey = apikey
	IdentifierType = identifiertype
	FilePath = filePath
	Route = route
	e.ApiKey = apikey
    e.IdentifierType = identifiertype
	e.FilePath = filePath
    
}

func SawoRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(Route, handler).Methods("GET")
	r.HandleFunc("/verify", payload_handler).Methods("POST")

	return r
}


func sawoSDK(w http.ResponseWriter, filename string, data interface{}) {
    t, err := template.ParseFiles(filename)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := t.Execute(w, data); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}



func handler(w http.ResponseWriter, r *http.Request) {
	
	configMap := map[string]interface{}{
		"apiKey": ApiKey, 
		"identifier_type": IdentifierType, 
		"file_path": FilePath}
	  fmt.Println(configMap)
	sawoSDK(w, FilePath, configMap)
}


func payload_handler(w http.ResponseWriter, r *http.Request) {
	 var data SawoPayload
   err := json.NewDecoder(r.Response.Body).Decode(&data)
   fmt.Println(err)
   if err != nil {
      http.Error(w, err.Error(), 500)
        return
   }
   
//    fmt.Println(data.UserID)
//    fmt.Println(data.VerificationToken)
}
