protos:
	buf generate

run:
	go run main.go

sample-request-observer:
	curl --request "POST" --header "Content-Type: application/json" --data '{}' http://localhost:12345/twirp/observer.Observer/Notify

sample-request-subject:
	curl --request "POST" --header "Content-Type: application/json" --data '{}' http://localhost:54321/twirp/subject.Subject/LatestState
