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



type SawoConfig struct{
	ApiKey string
	IdentifierType string
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

func (e *SawoConfig) Init(apikey string, identifiertype string) {
	ApiKey = apikey
	IdentifierType = identifiertype
	e.ApiKey = apikey
    e.IdentifierType = identifiertype
    
}

func SawoRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", main_handler).Methods("GET")
	r.HandleFunc("/login.html", handler).Methods("GET")
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
	
	configMap := map[string]interface{}{"apiKey": ApiKey, "identifier_type": IdentifierType}
	  fmt.Println(configMap)
	sawoSDK(w, "../assets/login.html", configMap)
}

func main_handler(w http.ResponseWriter, r *http.Request) {
	
	
	sawoSDK(w, "../assets/index.html", nil)
	
}

func payload_handler(w http.ResponseWriter, r *http.Request) {
	 var data SawoPayload
   err := json.NewDecoder(r.Response.Body).Decode(&data)
   fmt.Println(err)
   if err != nil {
      http.Error(w, err.Error(), 500)
        return
   }
   
   fmt.Println(data.UserID)
   fmt.Println(data.VerificationToken)
}