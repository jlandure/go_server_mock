/**
 * Created with IntelliJ IDEA.
 * User: julien_la
 * Date: 23/05/13
 * Time: 18:17
 * To change this template use File | Settings | File Templates.
 */
package optima

import (
	"fmt"
	"net/http"
	"encoding/xml"
	"encoding/json"
	"io/ioutil"
	//"bytes"
	"strings"
	"strconv"
	"log"
	"html/template"
	"time"
)


type OptimaResult struct {
	XMLName   			 xml.Name `xml:"http://sydev.com/ optima"`
	Error	  			int      `xml:"error"`
	ErrorMessage    	   string   `xml:"errorMessage,omitempty"`
	NumeroAffaireOptima    string 	 `xml:"numeroAffaireOptima"`
}

type IplusRequest struct {
	Uuid                              string  `json:"uuid"`
	IdOptima                          string   `json:"id_optima"`
	Montant                           float64 `json:"montant"`
	Avancement                        string  `json:"avancement"`
	//	Date_derniere_modification_optima uint64  `json:"date_derniere_modification_optima"`
	Date_derniere_modification_optima string  `json:"date_derniere_modification_optima"`
	//	Date_envoi_optima_iplus           uint64  `json:"date_envoi_optima_iplus"`
	Date_envoi_optima_iplus           string  `json:"date_envoi_optima_iplus"`

}

var Compteur int = 1200



func HandlerPostOptima(w http.ResponseWriter, r *http.Request) {
	log.Printf("Content-Type %s & Content-Length %d \n", r.Header.Get("Content-Type"), r.ContentLength)
	if r.Method == "POST" {
		login := r.FormValue("Login")
		log.Printf("Login %s\n", login)
		password := r.FormValue("Password")
		log.Printf("Password %s\n", password)
		uuid := r.FormValue("UUID")
		log.Printf("UUID %s\n", uuid)
		xmlString := r.FormValue("XMLOpportunity")
		log.Printf("XMLOpportunity %s\n", xmlString)

		var result *OptimaResult


		//for {
		//boucle infinie
		//	<-time.After(10 * time.Second)
		//		fmt.Printf("10s client waiting...\n")
		//	}
		//http.Error(w, "INTERNAL ERROR OPTIMA", http.StatusGatewayTimeout)


		Compteur++
		// write whole the body
		err := ioutil.WriteFile("Go_server_Optima/compteur.txt", []byte(strconv.Itoa(Compteur)), 0644)
		if err != nil { panic(err) }

		if (Compteur%3 == 0) {
			result = &OptimaResult{Error: 0, NumeroAffaireOptima: strconv.Itoa(Compteur) }
		} else if (Compteur%3 == 1) {
			result = &OptimaResult{Error: 1, ErrorMessage: "Error Msg : communication error with optima"}
		}

		//case if(compteur % 3 == 2)
		if (result == nil) {
			w.Header().Set("Content-Type", "application/xml")
			//resultFakeXml := []byte("<?xml version='1.0' encoding='utf-8'?><result><not_working value=\"0\"/></result>")
			resultFakeXml := []byte("toto</result>")
			w.Write(resultFakeXml)
			log.Printf("xml_data %s\n", resultFakeXml)
			return
		}
		if xml_data, err := xml.Marshal(result); err == nil {
			w.Header().Set("Content-Type", "application/xml")
			fmt.Fprintf(w, "<?xml version='1.0' encoding='utf-8'?>")
			w.Write(xml_data)
			log.Printf("xml_data %s\n", xml_data)
		} else {
			http.Error(w, "INTERNAL ERROR OPTIMA", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Not a POST Method", http.StatusNotFound)
		return
	}
}

func HandlerUpdateOpportunite(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Go_server_Optima/src/template/updateOpportuniteForm.html"))
	tc := make(map[string]interface{})
	tc["Title"] = "Formulaire de saisie du numero Affaire Optima"
	//	tc["Products"] = Products

	if err := tmpl.Execute(w, tc); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}



func HandlerUpdateOpportunitePost(w http.ResponseWriter, r *http.Request) {
	requestObject := IplusRequest{
		Uuid : r.FormValue("uuid"),
		IdOptima : r.FormValue("id_optima"),
		Avancement : r.FormValue("avancement"),
	}
	montant, _ := strconv.ParseFloat(r.FormValue("montant"), 32)
	requestObject.Montant = montant
	//test request Object AVancement
	//requestObject.Avancement = "TOTO";
	//requestObject.Date_derniere_modification_optima, _ = strconv.ParseUint(tranformDate(r.FormValue("date_derniere_modification_optima_date"), r.FormValue("date_derniere_modification_optima_time")),10,64)
	requestObject.Date_derniere_modification_optima = tranformDate(r.FormValue("date_derniere_modification_optima_date"), r.FormValue("date_derniere_modification_optima_time"))
	//requestObject.Date_envoi_optima_iplus, _ = strconv.ParseUint(tranformDate(r.FormValue("date_envoi_optima_iplus_date"), r.FormValue("date_envoi_optima_iplus_time")),10,64)
	requestObject.Date_envoi_optima_iplus = tranformDate(r.FormValue("date_envoi_optima_iplus_date"), r.FormValue("date_envoi_optima_iplus_time"))
	//no need : log.Print(requestObject);

	json_data, err := json.Marshal(requestObject);
	if (err != nil) {
		log.Fatal(err)
	}
	log.Printf("json_data %s\n", json_data)

	client := &http.Client{}

	//req, _ := http.NewRequest("POST", "http://localhost:9090/toto", strings.NewReader(string(json_data)))
	req, _ := http.NewRequest("POST", "http://localhost:8484/sfa-ws/rest/optima/updateOpportunite", strings.NewReader(string(json_data)))
	//http://localhost:8484/sfa-ws/rest/optima/updateOpportunite

	req.Header.Add("User-Agent", "Go-Client")
	req.Header.Add("Conf-Id", "INEO_ENERSYS");
	req.Header.Add("Username", "service.client@vecteurplus.com");
	req.Header.Add("Password", "admin");
	req.Header.Add("Content-Type", "application/json");
	resp, err := client.Do(req)
	if (err != nil) {
		log.Fatal(err)
	}

	var b []byte;
	b, err = ioutil.ReadAll(resp.Body);
	log.Printf("Error-Code : %s", resp.Header.Get("Error-Code"))
	log.Printf("Error-Message : %s", resp.Header.Get("Error-Message"))
	log.Printf("resultat body de la req : %s", string(b))
	http.Redirect(w, r, "/optima/updateOpportunite", http.StatusFound)
	defer resp.Body.Close()

}

const format string = "2006-01-02 15:04"
const formatIplus string = "200601021504"

func tranformDate(optimaDate string, optimaTime string) string {
	//log.Printf(" date %s, time :%s", optimaDate, optimaTime)
	dateTemp := optimaDate + " " + optimaTime
	//log.Printf("concat %s", dateTemp)
	realDate, err := time.Parse(format, dateTemp)
	if (err != nil) {
		log.Fatal(err)
	}
	//log.Print(realDate)
	dateToReturn := realDate.Format(formatIplus) + "00"
	//log.Print(dateToReturn)
	return dateToReturn
}
