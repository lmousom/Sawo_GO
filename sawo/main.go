package main

import (
	"net/http"

	sawo "github.com/latiful/sawogo"
)

func main(){

   sawoconfig := new(sawo.SawoConfig)
   // can be either 'email' or 'phone_number_sms'
   // add your api key here
   sawoconfig.Init("390508ed-d1c7-44a6-adbc-58ff2d33999e", "email")
   r := sawo.SawoRouter()
	http.ListenAndServe(":8080", r)


}