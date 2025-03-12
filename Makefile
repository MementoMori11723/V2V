run :
	@docker-compose -f config/compose.yml -p echo-flow build
	@docker-compose -f config/compose.yml -p echo-flow up -d
	@sleep 3
	@docker exec echo_ai bash -c "ollama pull qwen:0.5b"

stop :
	@docker-compose -f config/compose.yml -p echo-flow down --remove-orphans
