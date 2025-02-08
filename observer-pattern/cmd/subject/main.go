package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/antblood/observer-pattern/pkg/pb"
	"github.com/antblood/observer-pattern/pkg/subject"
)

func mockUpdateState(subject *subject.Subject) {
	i := 0
	for {
		time.Sleep(3 * time.Second)
		i++
		subject.UpdateState(context.Background(), &pb.UpdateStateRequest{
			State: fmt.Sprintf("updated %d", i),
		})
		subject.Broadcast(context.Background(), &pb.BroadcastRequest{})
	}
}

func main() {
	mux := http.NewServeMux()
	subject := subject.NewSubject([]string{"http://localhost:54321"})
	mux.Handle("/twirp/subject.Subject/", pb.NewSubjectServer(subject))
	go mockUpdateState(subject)
	http.ListenAndServe(":12345", mux)
}
