run-tests:
	go test -v
	(cd ./transports/console && go test -v)
