package main

import (
	"net/http"

	"github.com/antblood/observer-pattern/pkg/observer"
	pb "github.com/antblood/observer-pattern/pkg/pb"
)

func main() {
	mux := http.NewServeMux()
	observer := observer.NewObserver("http://localhost:12345")
	mux.Handle("/twirp/observer.Observer/", pb.NewObserverServer(observer))
	http.ListenAndServe(":54321", mux)
}
