
docker-up:
	docker compose up -d

migrate-dev:
	go run github.com/steebchen/prisma-client-go migrate dev --name add --preview-feature --create-only

migrate-deploy:
	go run github.com/steebchen/prisma-client-go migrate deploy --preview-feature