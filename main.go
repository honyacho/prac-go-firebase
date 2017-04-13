package main

import (
//    "fmt"
//    "net/http"
    "gopkg.in/zabawaba99/firego.v1"
    "golang.org/x/oauth2"
    "io/ioutil"
    "golang.org/x/oauth2/google"
    "log"
)

func main() {
    d, err := ioutil.ReadFile("id.json")
    if err != nil {
        return
    }

    conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
    if err != nil {
        return
    }
    fb := firego.New("https://test-slack-164207.firebaseio.com/", conf.Client(oauth2.NoContext))

    notifications := make(chan firego.Event)
    if err := fb.Watch(notifications); err != nil {
    	log.Fatal(err)
    }

    defer fb.StopWatching()
    for event := range notifications {
    	log.Printf("Event %#v\n", event)
    }

    log.Println("Notifications have stopped")
}

