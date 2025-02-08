package main

import (
	"context"
	"net/http"
	"time"

	pb "github.com/antblood/observer-pattern/pkg/pb"
	"github.com/antblood/observer-pattern/pkg/subject"
)

func main() {
	mux := http.NewServeMux()
	subject := subject.NewSubject([]pb.Observer{
		pb.NewObserverProtobufClient("http://localhost:54321", http.DefaultClient),
	})
	mux.Handle("/twirp/subject.Subject/", pb.NewSubjectServer(subject))
	go func() {
		for {
			time.Sleep(3 * time.Second)
			subject.UpdateState(context.Background(), &pb.UpdateStateRequest{
				State: "updated",
			})
			subject.Broadcast(context.Background(), &pb.BroadcastRequest{})
		}
	}()

	http.ListenAndServe(":12345", mux)
}
