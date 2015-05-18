# makefile for simple directory

neat:
	go fmt ./...

world:
	go build ./...

clean:
	go clean ./...
	find . -type f -name "*~" -delete
