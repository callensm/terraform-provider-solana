docs:
	tfplugindocs

clean:
	rm -rf vendor/

test:
	TF_ACC=1 go test -v -cover ./internal/...

vendor: clean
	go mod tidy && go mod vendor

.PHONY: docs clean test vendor
