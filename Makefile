default_goal: run

build:
	@echo "Building the program..."
	@go build -tags netgo -ldflags="-s -w" -o app .

run:
	@echo "Running the program..."
	if [ ! -f app ]; then make build; fi
	@./app

clean:
	@echo "Cleaning up..."
	@rm -f app

help:
	@echo "Usage: make [build|run|clean]"
	@echo "  build: Build the program"
	@echo "  run: Run the program"
	@echo "  clean: Clean up"

.PHONY: build run clean
