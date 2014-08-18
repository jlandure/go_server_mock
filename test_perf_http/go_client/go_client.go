package main

import (
	rest "github.com/ant0ine/go-json-rest"
	simplejson "github.com/bitly/go-simplejson"
	"net/http"
	"time"
	"encoding/json"
	"strings"
	"log"
	"io/ioutil"
)

type HttpCustomResponse struct {
	url      string
	response *http.Response
	err      error
}

type DataJsonForC struct {
	Call_a                         string `json:"call_a"`
	Call_b                         string `json:"call_b"`
}

type DataForC struct {
	ErrorForC                              string  `json:"error"`
	DataElementJsonForC 					   DataJsonForC  `json:"data"`
}

type DataToReturn struct {
	Id                    string
	Reference             string
	Type                  string
	Date_creation         string
	Date_modification     string
	Object_marche         string
}

const urlC = "http://localhost:9090/rest/c"
const urlD = "http://localhost:9090/rest/d?sort=a&filter=namedfilter&page=1&per_page=20"

func main() {

	handler := rest.ResourceHandler{
		EnableRelaxedContentType:true,
		DisableJsonIndent: true,
	}

	handler.SetRoutes(
		rest.Route{"GET", "/rest/api/v1/im", callScenario1},
	)
	http.ListenAndServe(":8080", &handler)
}

//func callD(url string, ch *chan) {
//	fmt.Printf("Calling D at %s", url)
//	resp, err := http.Get(url)
//	resp.Body.Close()
//	ch <- &HttpResponse{url, resp, err}
//
//}

func callScenario1(w *rest.ResponseWriter, r *rest.Request) {

	results := callCAndD();

	var returnedResultArray []DataToReturn
	for _, result := range results {
		log.Printf("Returned %s status: %s\n", result.url,
			result.response.Status)
		var jsonBody []byte;
		jsonBody, _ = ioutil.ReadAll(result.response.Body);
		log.Println("read all done len=>", len(jsonBody))
		simpleDataFromBody, err2 := simplejson.NewJson(jsonBody)
		log.Println("construction du json")
		if err2 != nil {
			log.Print("err2",err2)
		}
		//var returnedResult map[int]DataToReturn
		//returnedResultArray, err3 := simpleDataFromBody.Get("data").MustArray(returnedResultArray);
		//*if err3 != nil {
		//	log.Print(err3)
		//}
//		for i, v := range simpleDataFromBody.Get("data").MustArray() {
//				returnedResult[i], _ = v.(DataToReturn)
//				log.Println(i, v)
//		}

		if(simpleDataFromBody != nil && simpleDataFromBody.Get("data") == nil) {
			//cas ou on a directement un tableau
			log.Println("cas tableau")
			log.Printf("data : %s", simpleDataFromBody)
			elements, err3 := simpleDataFromBody.Array()
			if err3 != nil {log.Print(err3)}
			log.Println(elements)
		}
		//defer result.response.Body.Close()
	}

	w.WriteJson(returnedResultArray)

	log.Printf("Scenario 1 done")
}

func callCAndD() []*HttpCustomResponse {
	ch := make(chan *HttpCustomResponse, 2) // buffer size set to 2 calss
	responses := []*HttpCustomResponse{}
	go func(url string) {
		log.Printf("Fetching %s \n", url)
		jsonToSend := DataJsonForC{Call_a:"Json_DATA A",Call_b:"Json DATA B"}
		dataToSend := DataForC{ErrorForC:"Aucune", DataElementJsonForC:jsonToSend}

		log.Printf("data : %s",dataToSend)
		//var json_dataToSend []byte;
		json_dataToSend, err := json.Marshal(dataToSend);
		if (err != nil) {
			log.Print(err)
		}
		log.Printf("json_data %s\n", json_dataToSend)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", urlC, strings.NewReader(string(json_dataToSend)))
		req.Header.Add("X-Login", "service.client@vecteurplus.com")
		resp, err := client.Do(req)

		ch <- &HttpCustomResponse{url, resp, err}
	}(urlC)
	go func(url string) {
		log.Printf("Fetching %s \n", url)
		resp, err := http.Get(url)
		resp.Body.Close()
		ch <- &HttpCustomResponse{url, resp, err}
	}(urlD)

	for {
		select {
		case r := <-ch:
			log.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == 2 {
				log.Printf("All is done\n")
				return responses
			}
		case <-time.After(50*time.Millisecond):
			log.Printf(".")
		}
	}
	log.Printf("All is done2\n")
	return responses
}
