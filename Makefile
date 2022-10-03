test: 
	@go test -coverprofile cover.out
	
coverage: test
	@go tool cover -html=cover.out