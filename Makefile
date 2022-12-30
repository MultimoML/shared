# https://makefiletutorial.com/

.PHONY: traefik portainer datadog etcd

traefik:
	sudo docker-compose -f traefik/docker-compose.yml --env-file .env up -d --force-recreate

portainer:
	sudo docker-compose -f portainer/docker-compose.yml --env-file .env up -d --force-recreate

datadog:
	sudo docker-compose -f datadog/docker-compose.yml --env-file .env up -d --force-recreate

etcd:
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative etcd/main.proto
		cd etcd; go build .
		mv etcd/etcd .