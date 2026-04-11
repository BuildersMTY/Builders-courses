.PHONY: test
test:
	go test -run TestApplyMiddlewares -v -count=1 .
