package subject

import (
	"context"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct {
	url string
}

type Subject struct {
	observers []Observer
	state     string
}

func (s Subject) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	s.observers = append(s.observers, Observer{url: in.Url})
	return &pb.AddResponse{}, nil
}

func (s Subject) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	var tempObservers []Observer
	for i, o := range s.observers {
		if o.url != in.Url {
			tempObservers = append(tempObservers, s.observers[i])
		}
	}
	s.observers = tempObservers
	return &pb.RemoveResponse{}, nil
}

func (s Subject) LatestState(ctx context.Context, in *pb.LatestStateRequest) (*pb.LatestStateResponse, error) {
	return &pb.LatestStateResponse{State: s.state}, nil
}
