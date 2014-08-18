package main

import (
	rest "github.com/ant0ine/go-json-rest"
	simplejson "github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
	"time"
)

func main() {

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		DisableJsonIndent: true,
	}
	handler.SetRoutes(
		rest.Route{"POST", "/rest/c", postDataC},
		rest.Route{"GET", "/rest/d", getDataD},
		rest.Route{"PUT", "/rest/e/msg/:host", putDataE},
	)
	//TEST call to dataplus initData()
	//log.Printf("Started on 9090")
	http.ListenAndServe(":9090", &handler)

}

var CompteurD int = 0
var CompteurE int = 0

func initData() {

//	client := &http.Client{}
//	//req, _ := http.NewRequest("POST", "http://localhost:9090/toto", strings.NewReader(string(json_data)))
//	req, _ := http.NewRequest("GET", "http://dataplus-es1.cloudapp.net/dataplus4/im/_search", nil)
//
//	req.Header.Add("User-Agent", "Go-Client")
//	req.Header.Add("Content-Type", "application/json");
//	resp, err := client.Do(req)
//	if (err != nil) {
//		//log.Fatal(err)
//	}

	//var b []byte;
	//b, _ = ioutil.ReadAll(resp.Body);
	//if err != nil { _.Fatal(err) }
	//log.Printf("resultat body de la req : %s", string(b))

	//	simpleData, err := simplejson.NewJson(b)
	//	histData := simpleData.Get("hits")
	//	hitsArray := histData.Get("hits");
	//	firstIm := hitsArray.GetIndex(0)
	//	idIm := firstIm.Get("_source").Get("im")
	//	log.Printf("Data %s\n", idIm)

}

func postDataC(w *rest.ResponseWriter, r *rest.Request) {
	xlogin := r.Header.Get("X-Login")
	if (xlogin != "service.client@vecteurplus.com") {
		rest.Error(w, "X-Login header required", 400)
		return
	}
	var jsonBody []byte;
	jsonBody, _ = ioutil.ReadAll(r.Body);
	//if err != nil { err.Fatal(err) }
	simpleDataFromBody, err2 := simplejson.NewJson(jsonBody)
	if err2 != nil {
		rest.Error(w, "body (json) is empty", 400)
		return
	}
	//log.Printf("Data %s\n", simpleDataFromBody.Get("data").Get("call_a"))
	//log.Printf("Data %s\n", (simpleDataFromBody.Get("data").Get("call_a") == nil))
	if (simpleDataFromBody.Get("error") == nil || simpleDataFromBody.Get("error").MustString() != "Aucune") {
		rest.Error(w, "error is not valued to \"Aucune\"", 400)
		return
	}
	if (simpleDataFromBody.Get("data").Get("call_a").MustString() == "") {
		rest.Error(w, "call_a is not valued", 400)
		return
	}
	if (simpleDataFromBody.Get("data").Get("call_b").MustString() == "") {
		rest.Error(w, "call_b is not valued", 400)
		return
	}
	defer r.Body.Close()

	b := []byte(`{"result":"Success", "data" : [
    {
        "id": 11288051,
        "reference": "vp.pub.11288051",
        "type": "public",
        "date_creation": "2012-04-09T00:20:22+0000",
        "date_modification": "2012-04-16T11:51:24+0000",
        "objet_marche": "ma\u00EEtrise d'oeuvre pour l'installation d'une p\u00EAcherie d'anguille de d\u00E9valaison sur la S\u00E8vre niortaise."
    },
    {
        "id": 11288208,
        "reference": "vp.pub.11288208",
        "type": "public",
        "date_creation": "2012-04-09T03:17:43+0000",
        "date_modification": "2012-04-16T11:52:25+0000",
        "objet_marche": "Travaux d'am\u00E9nagement du lotissement ' La Grande Corv\u00E9e Pastey ' \u00E0 IZEURE (lot unique)"
    },
    {
        "id": 11288215,
        "reference": "vp.pub.11288215",
        "type": "public",
        "date_creation": "2012-04-09T03:24:13+0000",
        "date_modification": "2012-04-16T12:30:01+0000",
        "objet_marche": "EXTENSION DU RESEAU D'ASSAINISSEMENT -LIEU-DIT \" LE GRAND SERVIGNY \""
    },
    {
        "id": 11288227,
        "reference": "vp.pub.11288227",
        "type": "public",
        "date_creation": "2012-04-09T03: 34: 14+0000",
        "date_modification": "2012-04-16T13: 27: 39+0000",
        "objet_marche": "D\u00E9signationd'unmaitred'oeuvrepourlar\u00E9habilitationde30logementsIndividuels(18T3et12T4)et1logementindividuel(1T6)ensitehabite"
    },
    {
        "id": 11289698,
        "reference": "vp.pub.11289698",
        "type": "public",
        "date_creation": "2012-04-10T11: 04: 09+0000",
        "date_modification": "2012-04-16T15: 07: 16+0000",
        "objet_marche": "REAMENAGEMENTDEL'APPARTEMENTDELAPOSTE"
    },
    {
        "id": 11289737,
        "reference": "vp.pub.11289737",
        "type": "public",
        "date_creation": "2012-04-10T11: 12: 39+0000",
        "date_modification": "2012-04-16T14: 22: 27+0000",
        "objet_marche": "r\u00E9novationdelafacadearri\u00E8redu10b\u00E2timentscolairePr\u00E9vertII"

    },
    {
        "id": 11289751,
        "reference": "vp.pub.11289751",
        "type": "public",
        "date_creation": "2012-04-10T11: 20: 24+0000",
        "date_modification": "2012-04-16T15: 05: 48+0000",
        "objet_marche": "March\u00E9dema\u00EEtrised'oeuvrerelatif\u00E0l'isolationdesfacadesetsoustoituredelaDirectionD\u00E9partementaledelacoh\u00E9sionsocialeetdelaprotectiondespopulations."
    },
    {
        "id": 11289802,
        "reference": "vp.pub.11289802",
        "type": "public",
        "date_creation": "2012-04-10T11: 39: 40+0000",
        "date_modification": "2012-04-16T14: 21: 39+0000",
        "objet_marche": "Ma\u00EEtrised'oeuvrepourlareconstructiondelastationd'\u00E9puration"
    }
    ,
    {
        "id": 11289648,
        "reference": "vp.pub.11289648",
        "type": "public",
        "date_creation": "2012-04-10T10: 53: 07+0000",
        "date_modification": "2012-04-16T14: 39: 16+0000",
        "objet_marche": "Constructiond'uneSallePolyvalente"
    }
]}`)
	w.ResponseWriter.Header().Add("Content-Type", "application/json")
	w.ResponseWriter.Write(b)
}

func getDataD(w *rest.ResponseWriter, r *rest.Request) {
	//sort := r.PathParam("sort")
	queries := r.URL.Query()
	sort := queries.Get("sort")
	if (sort == "") {
		rest.Error(w, "sort required", 400)
		return
	}
	filter := queries.Get("filter")
	if (filter == "") {
		rest.Error(w, "filter required", 400)
		return
	}
	page := queries.Get("page")
	if (page == "") {
		rest.Error(w, "page required", 400)
		return
	}
	per_page := queries.Get("per_page")
	if (per_page == "") {
		rest.Error(w, "per_page required", 400)
		return
	}

	if (CompteurD%4 == 3) {
		w.ResponseWriter.WriteHeader(http.StatusNotModified)
		CompteurD++
		return
	}
	CompteurD++

	//	uriNext := r.UriForWithParams("rest/d", map[string][]string {"sort":{"a"}, "filter":{"namedfilter"}, "page":{"2"}, "per_page":{"20"}})
	//	uriPrevious := r.UriForWithParams("rest/d", map[string][]string {"sort":{"a"}, "filter":{"namedfilter"}, "page":{"1"}, "per_page":{"20"}})
	//	uriFirst := r.UriForWithParams("rest/d", map[string][]string {"sort":{"a"}, "filter":{"namedfilter"}, "page":{"1"}, "per_page":{"20"}})
	//	uriLast := r.UriForWithParams("rest/d", map[string][]string {"sort":{"a"}, "filter":{"namedfilter"}, "page":{"50"}, "per_page":{"20"}})
	//	link := fmt.Sprintf("%s; rel=\"next\", %s; rel=\"previous\", %s; rel=\"first\", %s; rel=\"last\"", uriNext, uriPrevious, uriFirst, uriLast)
	//	log.Printf("%s", uriNext)
	w.ResponseWriter.Header().Add("link", "<http://localhost:9090/rest/d?sort=a&filter=namedfilter&page=2&per_page=20>; rel=\"next\",<http://localhost:9090/rest/d?sort=a&filter=namedfilter&page=1&per_page=20>; rel=\"previous\", <http://localhost:9090/rest/d?sort=a&filter=namedfilter&page=1&per_page=100>; rel=\"first\", <http://localhost:9090/rest/d?sort=a&filter=namedfilter&page=50&per_page=20>; rel=\"last\"")
	//	w.ResponseWriter.Header().Add("link",link)


	b := []byte(`[
    {
        "id": 11288051,
        "reference": "vp.pub.11288051",
        "type": "public",
        "date_creation": "2012-04-09T00:20:22+0000",
        "date_modification": "2012-04-16T11:51:24+0000",
        "objet_marche": "ma\u00EEtrise d'oeuvre pour l'installation d'une p\u00EAcherie d'anguille de d\u00E9valaison sur la S\u00E8vre niortaise."
    },
    {
        "id": 11288208,
        "reference": "vp.pub.11288208",
        "type": "public",
        "date_creation": "2012-04-09T03:17:43+0000",
        "date_modification": "2012-04-16T11:52:25+0000",
        "objet_marche": "Travaux d'am\u00E9nagement du lotissement ' La Grande Corv\u00E9e Pastey ' \u00E0 IZEURE (lot unique)"
    },
    {
        "id": 11288215,
        "reference": "vp.pub.11288215",
        "type": "public",
        "date_creation": "2012-04-09T03:24:13+0000",
        "date_modification": "2012-04-16T12:30:01+0000",
        "objet_marche": "EXTENSION DU RESEAU D'ASSAINISSEMENT -LIEU-DIT \" LE GRAND SERVIGNY \""
    },
    {
        "id": 11288227,
        "reference": "vp.pub.11288227",
        "type": "public",
        "date_creation": "2012-04-09T03: 34: 14+0000",
        "date_modification": "2012-04-16T13: 27: 39+0000",
        "objet_marche": "D\u00E9signationd'unmaitred'oeuvrepourlar\u00E9habilitationde30logementsIndividuels(18T3et12T4)et1logementindividuel(1T6)ensitehabite"
    },
    {
        "id": 11289698,
        "reference": "vp.pub.11289698",
        "type": "public",
        "date_creation": "2012-04-10T11: 04: 09+0000",
        "date_modification": "2012-04-16T15: 07: 16+0000",
        "objet_marche": "REAMENAGEMENTDEL'APPARTEMENTDELAPOSTE"
    },
    {
        "id": 11289406,
        "reference": "vp.pub.11289406",
        "type": "public",
        "date_creation": "2012-04-10T11: 12: 39+0000",
        "date_modification": "2012-04-16T14: 22: 27+0000",
        "objet_marche": "Mocking r\u00E9novationdelafacadearri\u00E8redu10b\u00E2timentscolairePr\u00E9vertII"

    },
    {
        "id": 11289407,
        "reference": "vp.pub.11289407",
        "type": "public",
        "date_creation": "2012-04-10T11: 20: 24+0000",
        "date_modification": "2012-04-16T15: 05: 48+0000",
        "objet_marche": "Mocking March\u00E9dema\u00EEtrised'oeuvrerelatif\u00E0l'isolationdesfacadesetsoustoituredelaDirectionD\u00E9partementaledelacoh\u00E9sionsocialeetdelaprotectiondespopulations."
    },
    {
        "id": 11289408,
        "reference": "vp.pub.11289408",
        "type": "public",
        "date_creation": "2012-04-10T11: 39: 40+0000",
        "date_modification": "2012-04-16T14: 21: 39+0000",
        "objet_marche": "Mocking Ma\u00EEtrised'oeuvrepourlareconstructiondelastationd'\u00E9puration"
    }
    ,
    {
        "id": 11289409,
        "reference": "vp.pub.11289409",
        "type": "public",
        "date_creation": "2012-04-10T10: 53: 07+0000",
        "date_modification": "2012-04-16T14: 39: 16+0000",
        "objet_marche": "Mocking Constructiond'uneSallePolyvalente"
    }
]`)
	w.ResponseWriter.Header().Add("Content-Type", "application/json")
	w.ResponseWriter.Write(b)
}

func putDataE(w *rest.ResponseWriter, r *rest.Request) {
	host := r.PathParam("host")
	if (host == "") {
		rest.Error(w, "host is empty", 400)
		return
	}
	var jsonBody []byte;
	jsonBody, _ = ioutil.ReadAll(r.Body);
	//if err != nil { log.Fatal(err) }
	simpleDataFromBody, err2 := simplejson.NewJson(jsonBody)
	if err2 != nil {
		rest.Error(w, "body (json) is empty", 400)
		return
	}

	if (simpleDataFromBody.Get("message").MustString() == "") {
		rest.Error(w, "message for " + host + " is not valued", 400)
		return
	}
	if (simpleDataFromBody.Get("details").MustString() == "") {
		rest.Error(w, "details for " + host + " is not valued", 400)
		return
	}

	CompteurE++
	if(CompteurE%2 == 0) {
		time.Sleep(500 * time.Millisecond)
	}

	defer r.Body.Close()
	w.ResponseWriter.Header().Add("X-Host", host)
}
