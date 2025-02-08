package observer

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct {
	subjectClient pb.Subject
}

func NewObserver(subjectUrl string) *Observer {
	subjectClient := pb.NewSubjectProtobufClient(subjectUrl, http.DefaultClient)
	return &Observer{subjectClient: subjectClient}
}

func (o Observer) Notify(ctx context.Context, in *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	fmt.Println("Observer: Received state change")
	newState, err := o.subjectClient.LatestState(ctx, &pb.LatestStateRequest{})
	if err != nil {
		fmt.Println("Observer: Error getting new state", err)
	}
	fmt.Println("Observer: New state", newState)
	return &pb.NotifyResponse{}, nil
}
