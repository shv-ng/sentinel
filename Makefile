up:
	docker-compose -f ./deploy/docker-compose.yml up -d

down:
	docker-compose -f ./deploy/docker-compose.yml down

redis-cli:
	docker-compose -f ./deploy/docker-compose.yml exec redis redis-cli
