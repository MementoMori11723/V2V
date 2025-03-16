run :
	@docker-compose -f config/compose.yml -p echo-flow build
	@docker-compose -f config/compose.yml -p echo-flow up -d

stop :
	@docker-compose -f config/compose.yml -p echo-flow down --remove-orphans
