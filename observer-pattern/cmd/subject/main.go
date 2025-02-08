package main

import (
	"net/http"

	pb "github.com/antblood/observer-pattern/pkg/pb"
	"github.com/antblood/observer-pattern/pkg/subject"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/twirp/subject.Subject/", pb.NewSubjectServer(&subject.Subject{}))

	http.ListenAndServe(":12345", mux)
}
