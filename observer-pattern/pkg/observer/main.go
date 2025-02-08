package observer

import (
	"context"
	"fmt"

	pb "github.com/antblood/observer-pattern/pkg/pb"
)

type Observer struct{}

func (o Observer) Notify(ctx context.Context, in *pb.NotifyRequest) (*pb.NotifyResponse, error) {
	fmt.Println("Observer: Received state change")
	return &pb.NotifyResponse{}, nil
}
