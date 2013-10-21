/**
 * Created with IntelliJ IDEA.
 * User: julien_la
 * Date: 23/05/13
 * Time: 09:55
 * To change this template use File | Settings | File Templates.
 */
package mercure

import (
	"net/http"
	"io/ioutil"
	"log"
)

func HandlerMercureSessionPost(w http.ResponseWriter, r *http.Request) {
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

		resultFakeXml := []byte("<LoginResult>golang</LoginResult>")
		w.Write(resultFakeXml)
		log.Printf("xml_data %s\n", resultFakeXml)
		return
	}
}

func HandlerMercureQueryPost(w http.ResponseWriter, r *http.Request) {
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
<ExecuteSqlResult><fsQuery><Result><Flds id="0"><CpyAddr1StrName Val="BP 113"></CpyAddr1StrName><CpyAddr1Complt Val="10 AVENUE REIMS"></CpyAddr1Complt><CpyAddr2Complt></CpyAddr2Complt><CpyType Val="Organisme Public"></CpyType><CpyTxtOpnFld11 Val="Non"></CpyTxtOpnFld11><CpyAddr1Postcode Val="44000"></CpyAddr1Postcode><CpyTxtOpnFld29></CpyTxtOpnFld29><CpyTxtOpnFld1 Val="952333"></CpyTxtOpnFld1><CpyTxtOpnFld2 Val=""></CpyTxtOpnFld2><CpyTxtOpnFld7 Val="Enseignement"></CpyTxtOpnFld7><CpyTxtOpnFld6 Val="85.32Z"></CpyTxtOpnFld6><CpyAddr1Country></CpyAddr1Country><CpyName Val="GOLANG COMPANY"></CpyName><CpyTxtOpnFld9 Val="Administrations - Collectivités"></CpyTxtOpnFld9><CpyTxtOpnFld3 Val=""></CpyTxtOpnFld3><CpyTxtOpnFld5 Val="19600049100099"></CpyTxtOpnFld5><CpyTxtOpnFld35 Val="Client"></CpyTxtOpnFld35><CpyTxtOpnFld8 Val="Enseignement"></CpyTxtOpnFld8><CpyAddr1City Val="Nantes"></CpyAddr1City></Flds></Result></fsQuery></ExecuteSqlResult>
		`)
		w.Write(resultFakeXml)
		log.Printf("xml_data %s\n", resultFakeXml)
		return
	}
}


func HandlerMercureScriptPost(w http.ResponseWriter, r *http.Request) {
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
<ExecuteServerScriptResult>123456-go</ExecuteServerScriptResult>
				`)
		w.Write(resultFakeXml)
		log.Printf("xml_data %s\n", resultFakeXml)
		return
	}
}
