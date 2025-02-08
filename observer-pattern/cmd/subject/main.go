package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	pb "github.com/antblood/observer-pattern/pkg/pb"
	"github.com/antblood/observer-pattern/pkg/subject"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/twirp/subject.Subject/", pb.NewSubjectServer(&subject.Subject{
		State: "initial",
		Observers: []pb.Observer{
			pb.NewObserverProtobufClient("localhost:54321", http.DefaultClient),
		},
	}))
	go func() {
		for {
			time.Sleep(3 * time.Second)
			http.Post(
				fmt.Sprintf("http://localhost:12345/twirp/subject.Subject/UpdateState"),
				"application/json",
				bytes.NewBuffer([]byte(`{"state": "updated"}`)),
			)
			http.Post(
				fmt.Sprintf("http://localhost:12345/twirp/subject.Subject/Broadcast"),
				"application/json",
				bytes.NewBuffer([]byte(`{}`)),
			)
		}
	}()

	http.ListenAndServe(":12345", mux)
}
