.PHONY: generate run build clean

generate:
	templ generate

run: generate
	cd cmd/server && go run main.go

build: generate
	go build -o bin/server cmd/server/main.go

clean:
	rm -rf bin/
	find . -name "*_templ.go" -type f -delete
