/**
 * Created with IntelliJ IDEA.
 * User: julien_la
 * Date: 26/04/13
 * Time: 09:26
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"log"
	"mercure"
	"eiffagetp"
	"optima"
)


func handler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body);
	if (err != nil) {
		log.Fatal("error handler: ", err)
	}
	//dont' write until you have read all the elements in the request
	fmt.Fprintf(w, "Hi there, I love Go Lang !\n")
	fmt.Fprintf(w, "User-Agent %s\n", r.UserAgent())
	fmt.Fprintf(w, "Body %s\n", string(b))
	defer r.Body.Close()
}

func main() {
	// read whole the file
	b, err := ioutil.ReadFile("Go_server_Optima/compteur.txt")
	//if err != nil { panic(err) }
	if err == nil {
		optima.Compteur, err = strconv.Atoi(string(b))
	}

	http.HandleFunc("/optima/PushOpportunity", optima.HandlerPostOptima)
	http.HandleFunc("/optima/updateOpportunite", optima.HandlerUpdateOpportunite)
	http.HandleFunc("/optima/updateOpportunite/post", optima.HandlerUpdateOpportunitePost)
	http.HandleFunc("/mercure/session", mercure.HandlerMercureSessionPost)
	http.HandleFunc("/mercure/query", mercure.HandlerMercureQueryPost)
	http.HandleFunc("/mercure/script", mercure.HandlerMercureScriptPost)
	http.HandleFunc("/eiffagetp/session", eiffagetp.HandlerEiffageSessionPost)
	http.HandleFunc("/eiffagetp/query", eiffagetp.HandlerEiffageQueryPost)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}


