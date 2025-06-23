.DEFAULT_GOAL := generate

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

up:
	go get -u ./... && go mod tidy
	go get -u -t ./... && go mod tidy
	cd examples && go get -u ./... && go mod tidy
	cd examples && go get -u -t ./... && go mod tidy

prepare:
	go mod tidy

commit: generate
	git add --all defaulttints/* *.gen.go DEFAULT_TINTS.md
	git commit -m "chore(tints): generate updated tints"

generate: license prepare
	mkdir -p defaulttints
	rm -rf defaulttints/* *.gen.go DEFAULT_TINTS.md
	go run github.com/lrstanley/bubbletint/cmd/tintgen
	gofmt -e -s -w defaulttints/*.go
	gofmt -e -s -w *.go
	go vet defaulttints/*.go
	go vet *.go
	go test -v ./...
