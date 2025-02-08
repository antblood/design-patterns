package main

import (
	"net/http"

	"github.com/antblood/observer-pattern/pkg/subject"
	pb "github.com/antblood/observer-pattern/pkg/pb"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/twirp/subject.Subject/", pb.NewSubjectServer(&subject.Subject{}))

	http.ListenAndServe(":12345", mux)
}
