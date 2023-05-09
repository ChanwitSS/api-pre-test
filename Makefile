run:
	swag init
	go run main.go
run-dev:
	swag init
	air
test:
	go test -v post/testing
clean:
	rm -rf post
	go clean -i .