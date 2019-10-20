.PHONY : build-osx build-linux test format gen-proto

format:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

build-osx:
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o bin/web cmd/web/web.go

build-linux: 
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o bin/web cmd/web/web.go

gen-proto:
	protoc --go_out=plugins=grpc:. protos/*.proto

update-bin:
	cp ../web ./bin/ && chmod +x ./bin/web