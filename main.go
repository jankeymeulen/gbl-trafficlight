package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"google.golang.org/appengine"
	"time"
	"google.golang.org/appengine/datastore"
	"regexp"
	//"google.golang.org/appengine/log"
)

type Effect struct {
	Name	string
	Colour	Colour
}

type DynamiteReply struct {
	Text	string	`json:"text"`
}

type DynamiteCall struct {
	Type      string    `json:"type"`
	Token     string    `json:"token"`
	EventTime time.Time `json:"event_time"`
	Space     struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Type        string `json:"type"`
	} `json:"space"`
	Message struct {
		Name   string `json:"name"`
		Sender struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
			AvatarURL   string `json:"avatarUrl"`
			Email       string `json:"email"`
		} `json:"sender"`
		CreateTime time.Time `json:"createTime"`
		Text       string    `json:"text"`
		Thread     struct {
			Name string `json:"name"`
		} `json:"thread"`
	} `json:"message"`
}

func dynamiteHandler(w http.ResponseWriter, r *http.Request) {

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
	
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx,"DynamiteCall",nil)
	if _,err := datastore.Put(ctx,key,&call) ; err != nil {
		http.Error(w,"Error in datastore: " + err.Error(), 400)
		return
	}
	
	var reply DynamiteReply

	if call.Type == "ADDED_TO_SPACE" {
		reply.Text = "Thanks for adding me! You can control me by telling me which effect "+
		"to display and in which colour. For example: \"SOLID 255 0 0\" turns me solid red."
		}
	if call.Type == "MESSAGE" {
			var response = handleMessage(call)
			reply.Text = "Message received! " + response
		}
	w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(http.StatusCreated)
    	json.NewEncoder(w).Encode(reply)
}

func handleMessage (call DynamiteCall) (string) {
	re := regexp.MustCompile("^((SOLID)|(BLINK)) [[:digit:]]{1,3} [[:digit:]]{1,3} [[:digit:]]{1,3}$")
	if re.MatchString(call.Message.Text) {
		return "Valid command: " + call.Message.Text
	} else {
		return "Invalid command: " + call.Message.Text
	}
}

func parseEffect (message string) (Effect) {
	var e Effect
	e.Name = message
	return e
}

func effectHandler(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		q := datastore.NewQuery("DynamiteCall").Order("-EventTime").Limit(1)
		var calls []DynamiteCall
		if _, err := q.GetAll(ctx, &calls); err != nil {
			http.Error(w,"Error retrieving from datastore: " + err.Error(), 400)
			return
		}
		var e Effect = parseEffect(calls[0].Message.Text)
		fmt.Fprintf(w,e.Name) //+","+e.Colour.R+","+e.Colour.G+","+e.Colour.B)
}

func main() {
	http.HandleFunc("/dynamite",dynamiteHandler)
	http.HandleFunc("/effect",effectHandler)
	appengine.Main()
}
