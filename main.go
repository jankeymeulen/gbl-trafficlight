package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"time"
	//"google.golang.org/appengine/datastore"
	//"google.golang.org/appengine/log"
)

type Pattern struct {
	Name 	string
	R	int
	G	int
	B	int
	Datestamp	time.Time
}

type DynamiteSender struct {
	name	string
	email	string
}

type DynamiteMessage struct {
	text	string
	sender	DynamiteSender
}

type DynamiteCall struct {
	calltype	string `json:"type"`
	message 	DynamiteMessage
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var call DynamiteCall

	if r.Body == nil {
		http.Error(w,"Missing request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&call)
	if err != nil {
		http.Error(w,"Error parsing request: " + err.Error(),400)
			return
		}	
	fmt.Fprintln(w,"Hello world! "+call.message)
}

func main() {
	http.HandleFunc("/push",indexHandler)
	appengine.Main()
}
