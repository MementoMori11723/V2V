default_goal: run

build:
	@echo "Running the program..."
	@go build -tags netgo -ldflags="-s -w" -o app .

run:
	@./app
