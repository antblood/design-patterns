package main

import (
	"net/http"

	"github.com/antblood/observer-pattern/pkg/observer"
	pb "github.com/antblood/observer-pattern/pkg/pb"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/twirp/observer.Observer/", pb.NewObserverServer(&observer.Observer{
		SubjectUrl: "localhost:12345",
	}))

	http.ListenAndServe(":54321", mux)
}
