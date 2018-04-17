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

type Effect struct {
	Name	string
	Colour struct {
		R int8
		G int8
		B int8
	}
	UpdateTime	time.Time
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
	
	var reply DynamiteReply

	if call.Type == "ADDED_TO_SPACE" {
			reply.Text = "Thanks for adding me!"
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
	return call.Message.Text
}

func effectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello world!")
}

func main() {
	http.HandleFunc("/dynamite",dynamiteHandler)
	http.HandleFunc("/effect",effectHandler)
	appengine.Main()
}
