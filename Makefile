default_goal: run

build:
	@echo "Running the program..."
	@go build -ldflags="-s -w" -o app .

run: build
	@./app
