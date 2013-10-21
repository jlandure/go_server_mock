/**
 * Created with IntelliJ IDEA.
 * User: julien_la
 * Date: 23/05/13
 * Time: 18:21
 * To change this template use File | Settings | File Templates.
 */
package eiffagetp

import (
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

func HandlerEiffageSessionPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Content-Type %s & Content-Length %d \n", r.Header.Get("Content-Type"), r.ContentLength)
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body);
		if (err != nil) {
			log.Fatal("error handler: ", err)
		}
		//dont' write until you have read all the elements in the request
		log.Printf("User-Agent %s\n", r.UserAgent())
		log.Printf("Body %s\n", string(b))
		defer r.Body.Close()

		resultFakeXml := []byte(`
<Result><sessionId>golang-eiffage-session</sessionId><serverUrl>http://INFO-195:9090/eiffagetp/query</serverUrl></Result>
				`)
		w.Write(resultFakeXml)
		log.Printf("xml_data %s\n", resultFakeXml)
		return
	}
}

func HandlerEiffageQueryPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("Content-Type %s & Content-Length %d \n", r.Header.Get("Content-Type"), r.ContentLength)
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body);
		if (err != nil) {
			log.Fatal("error handler: ", err)
		}
		//dont' write until you have read all the elements in the request
		log.Printf("User-Agent %s\n", r.UserAgent())
		log.Printf("Body %s\n", string(b))
		defer r.Body.Close()

		var resultFakeXml []byte
		if strings.Contains(string(b), "FROM Opportunity") {
			resultFakeXml = []byte(`
<result><records><sf:Id>123456</sf:Id><sf:OpportunityNumber__c>789012</sf:OpportunityNumber__c><sf:VP_ID__c>C3587C35-AFE8-4434-AD69-104C84025163</sf:VP_ID__c><sf:Name>Golang Salesforce Affaire</sf:Name></records></result>
				`)
		} else if strings.Contains(string(b), "FROM Ville__c") {
			resultFakeXml = []byte(`
<result><records><sf:Id>555666</sf:Id><sf:Name>NANTES</sf:Name><sf:Departement__c>44 - Loire-Atlantique</sf:Departement__c></records></result>
				`)
		}

		w.Write(resultFakeXml)
		log.Printf("xml_data %s\n", resultFakeXml)
		return
	}
}


