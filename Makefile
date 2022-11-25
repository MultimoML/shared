# https://makefiletutorial.com/

.PHONY: traefik portainer

traefik:
	sudo podman-compose up -d -f ./traefik/docker-compose.yml --env-file .env --force-recreate

portainer:
	sudo podman-compose up -d -f ./portainer/docker-compose.yml --env-file .env --force-recreate
