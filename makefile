test:
	go test -coverprofile fmtcoverage.out ./...
	go tool cover -func=fmtcoverage.out