run: build start

build: # Build a binary
	@echo "  >  Building binary..."
	go build -o ./location ./cmd/main.go

start: # Start an application
	@echo "  >  Starting binary..."
	@./location

test:
	go test -bench=. ./... -benchmem
