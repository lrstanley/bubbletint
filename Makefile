.DEFAULT_GOAL := generate

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

up:
	cd ./cmd/tintgen && go get -u -t ./... && go mod tidy
	cd ./examples && go get -u -t ./... && go mod tidy
	go get -u -t ./... && go mod tidy

prepare:
	cd ./cmd/tintgen && go mod tidy
	cd ./examples && go mod tidy
	go mod tidy

commit: generate
	git add --all defaulttints/* *.gen.go
	git commit -m "chore(tints): generate updated tints"

generate: license prepare
	rm -rf ./public
	cd ./cmd/tintgen && go run . ../../
	gofmt -e -s -w *.go
	go vet *.go
	go test -v ./...
