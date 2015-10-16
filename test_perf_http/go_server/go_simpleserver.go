package main

import (
	rest "github.com/ant0ine/go-json-rest/rest"
	//simplejson "github.com/bitly/go-simplejson"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, _ := rest.MakeRouter(
		rest.Get("/rest/test", test),
	)
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

}

type RefType struct {
	Id        string `json:"id"`
	Reference string `json:"dataref"`
	Kind      string `json:"kind"`
}

func test(w rest.ResponseWriter, r *rest.Request) {
	queries := r.URL.Query()
	idQuery := queries.Get("id")
	if idQuery == "" {
		rest.Error(w, "idQuery required", 400)
		return
	}

	ch := make(chan *RefType, 2) // buffer size set to 2 channels
	responses := []*RefType{}

	go func(idQuery string) {
		log.Printf("Fetching routine1 %s \n", idQuery)

		ch <- &RefType{Id: idQuery, Reference: "toto", Kind: "offre"}
	}(idQuery)
	go func(idQuery string) {
		log.Printf("Fetching routine2 %s \n", idQuery)
		ch <- &RefType{Id: idQuery, Reference: "tata", Kind: "option"}
	}("6")

	continu := true
	for continu {
		select {
		case r := <-ch:
			log.Printf("%s was fetched\n", r.Id)
			responses = append(responses, r)
			if len(responses) == 2 {
				log.Printf("Routines is done\n")
				continu = false
			}
		case <-time.After(50 * time.Millisecond):
			log.Printf(".")
		}
	}
	log.Printf("All is done\n")
	w.Header().Set("Content-Type", "application/json")
	dataToSend, _ := json.Marshal(responses)

	w.(http.ResponseWriter).Write(dataToSend)
	//w.(http.ResponseWriter).Write([]byte(`{"result":"Success"}`))
}
