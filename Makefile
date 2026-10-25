.DEFAULT_GOAL := generate

export ITERM_SCHEMES_VERSION := release-20261019-151001-86a776d

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

up:
	cd ./cmd/tintgen && go get -u -t ./... && go mod tidy
	cd ./glamourtint && go get -u -t ./... && go mod tidy
	cd ./_examples && go get -u -t ./... && go mod tidy
	go get -u -t ./... && go mod tidy

prepare:
	cd ./cmd/tintgen && go mod tidy
	cd ./glamourtint && go mod tidy
	cd ./_examples && go mod tidy
	go mod tidy

generate: license prepare
	rm -rf ./public
	cd ./cmd/tintgen && go test ./...
	cd ./cmd/tintgen && go run . ../../
	gofmt -e -s -w *.go
	go vet *.go
	go test -v ./...
	cd ./glamourtint && go test ./...
