package subject

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct {
	Url string
}

type Subject struct {
	Observers []Observer
	State     string
}

func (s Subject) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	s.Observers = append(s.Observers, Observer{Url: "in.Url"})
	return &pb.AddResponse{}, nil
}

func (s Subject) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	var tempObservers []Observer
	for i, o := range s.Observers {
		if o.Url != "in.Url" {
			tempObservers = append(tempObservers, s.Observers[i])
		}
	}
	s.Observers = tempObservers
	return &pb.RemoveResponse{}, nil
}

func (s Subject) LatestState(ctx context.Context, in *pb.LatestStateRequest) (*pb.LatestStateResponse, error) {
	return &pb.LatestStateResponse{
		State: s.State,
	}, nil
}

func (s *Subject) UpdateState(ctx context.Context, in *pb.UpdateStateRequest) (*pb.UpdateStateResponse, error) {
	s.State = in.State
	return &pb.UpdateStateResponse{}, nil
}

func (s *Subject) Broadcast(ctx context.Context, in *pb.BroadcastRequest) (*pb.BroadcastResponse, error) {
	fmt.Println("broadcasting")
	for _, o := range s.Observers {
		fmt.Println("broadcasting to", o.Url)
		resp, err := http.Post(fmt.Sprintf("http://%s/twirp/observer.Observer/Notify", o.Url), "application/json", bytes.NewBuffer([]byte(`{}`)))
		if err != nil {
			fmt.Println("error broadcasting to", o.Url, err)
		}
		fmt.Println("response", resp)
	}
	return &pb.BroadcastResponse{}, nil
}
