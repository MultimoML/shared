# https://makefiletutorial.com/

.PHONY: traefik portainer

traefik:
	sudo docker-compose -f traefik/docker-compose.yml --env-file .env up -d --force-recreate

portainer:
	sudo docker-compose -f portainer/docker-compose.yml --env-file .env up -d --force-recreate
