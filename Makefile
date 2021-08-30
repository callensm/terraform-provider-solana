clean:
	rm -rf vendor/

test:
	TF_ACC=1 go test -v -cover ./solana/...

vendor: clean
	go mod tidy && go mod vendor

.PHONY: clean test vendor
