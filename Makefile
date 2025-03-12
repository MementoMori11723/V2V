run :
	@docker-compose -f config/compose.yml -p echo-flow build
	@docker-compose -f config/compose.yml -p echo-flow up -d
	@sleep 3
	@docker exec echo_ai bash -c "ollama pull smollm:135m"

stop :
	@docker-compose -f config/compose.yml -p echo-flow down --remove-orphans
