package observer

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct {
	SubjectUrl string
}

func (o Observer) Notify(ctx context.Context, in *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	fmt.Println("Observer: Received state change")
	newState, err := http.Post(fmt.Sprintf("http://%s/twirp/subject.Subject/LatestState", o.SubjectUrl), "application/json", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		fmt.Println("Observer: Error getting new state", err)
	}
	fmt.Println("Observer: New state", newState)
	return &pb.NotifyResponse{}, nil
}
