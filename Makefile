run :
	@docker-compose -f config/compose.yml build
	@docker-compose -f config/compose.yml up -d
	@docker exec echo_ai bash -c "ollama pull qwen2.5:1.5b"

stop :
	@docker-compose -f config/compose.yml down --remove-orphans
