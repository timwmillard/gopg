
-include .env

DATABASE_URL=postgresql://$(PGUSER):$(PGPASSWORD)@$(PGHOST):$(PGPORT)/$(PGDATABASE)

env:
	@echo ENV=$(ENV)
	@echo PORT=$(PORT)
	@echo PGHOST=$(PGHOST)
	@echo PGPORT=$(PGPORT)
	@echo PGDATABASE=$(PGDATABASE)
	@echo PGUSER=$(PGUSER)
	@echo PGPASSWORD=$(PGPASSWORD)
	@echo DATABASE_URL=$(DATABASE_URL)

docker-db: ## Start the database using docker
	docker start $(PGDATABASE) || \
	docker run --name $(PGDATABASE) \
		-p $(PGPORT):5432 \
		-e POSTGRES_USER=$(PGUSER)  \
		-e POSTGRES_PASSWORD=$(PGPASSWORD) \
		-e POSTGRES_DB=$(PGDATABASE) \
		-d postgres:15-alpine

docker-psql: ## Connect to psql running in the docker container
	docker exec -it $(PGDATABASE) psql -U $(PGUSER) -d $(PGDATABASE)

docker-logs: ## Show the docker database logs
	docker logs $(PGDATABASE) -f

docker-clean:  ## Remove the docker database and container
	-docker stop $(PGDATABASE)
	-docker rm $(PGDATABASE)

