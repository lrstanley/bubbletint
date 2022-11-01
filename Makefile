.DEFAULT_GOAL := generate

license:
	curl -sL https://liam.sh/-/gh/g/license-header.sh | bash -s

up: go-upgrade-deps
	@echo

go-fetch:
	go mod download
	go mod tidy

go-upgrade-deps:
	go get -u ./...
	go mod tidy

go-upgrade-deps-patch:
	go get -u=patch ./...
	go mod tidy

commit:
	git add --all tints/* *.gen.go
	git commit -m "chore(tints): generate updated tints"

generate:
	mkdir -p tints
	rm -rf tints/* *.gen.go
	# go run cmd/tintgen/*.go
	go run github.com/lrstanley/bubbletint/cmd/tintgen
	gofmt -e -s -w tints/*.go
	gofmt -e -s -w *.go
	go vet tints/*.go
	go vet *.go
