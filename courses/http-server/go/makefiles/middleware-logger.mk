.PHONY: test
test:
	go build -o /tmp/http-server . && exec /tmp/http-server
