package subject

import (
	"context"
	"fmt"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct {
	Url string
}

type Subject struct {
	Observers []pb.Observer
	State     string
}

func (s Subject) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	// logic to add observer
	return &pb.AddResponse{}, nil
}

func (s Subject) Remove(ctx context.Context, in *pb.RemoveRequest) (*pb.RemoveResponse, error) {
	// logic to remove observer
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
		resp, err := o.Notify(ctx, &pb.NotifyRequest{})
		if err != nil {
			fmt.Println("error broadcasting to", o, err)
		}
		fmt.Println("response", resp)
	}
	return &pb.BroadcastResponse{}, nil
}
