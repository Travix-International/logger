run-tests:
	go test -cover -v

cover:
	go test -coverprofile=cover.tmp && go tool cover -html=cover.tmp
