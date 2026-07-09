.DEFAULT_GOAL := build

# Update GoFrame and its CLI to latest stable version.
.PHONY: up
up: cli.install
	@gf up -a

# Build binary using configuration from hack/config.yaml.
.PHONY: build
build: cli.install
	@gf build -ew

# Parse api and generate controller/sdk.
.PHONY: ctrl
ctrl: cli.install
	@gf gen ctrl

# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao

# Run database migrations (goose).
GOOSE_DIR     = ./manifest/migration
GOOSE_DB      = "root:@tcp(127.0.0.1:3306)/go_demo?parseTime=true&multiStatements=true"

.PHONY: migrate
migrate: goose.install
	@goose -dir $(GOOSE_DIR) mysql $(GOOSE_DB) up

.PHONY: migrate-down
migrate-down: goose.install
	@goose -dir $(GOOSE_DIR) mysql $(GOOSE_DB) down

.PHONY: migrate-status
migrate-status: goose.install
	@goose -dir $(GOOSE_DIR) mysql $(GOOSE_DB) status

.PHONY: migrate-create
migrate-create: goose.install
	@test -n "$(NAME)" || (echo "用法: make migrate-create NAME=add_xxx_table" && exit 1)
	@goose -dir $(GOOSE_DIR) create $(NAME) sql

.PHONY: goose.install
goose.install:
	@set -e; \
	goose -version > /dev/null 2>&1 || (echo "安装 goose..." && go install github.com/pressly/goose/v3/cmd/goose@latest)

# Parse current project go files and generate enums go file.
.PHONY: enums
enums: cli.install
	@gf gen enums

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@gf gen service


# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git rev-parse --short HEAD))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@gf docker ${_PUSH} -tn $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push: cli.install
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy: cli.install
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_ENV};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	if [ $(DEPLOY_NAME) != "" ]; then \
		kubectl patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}"; \
	fi;


# Parsing protobuf files and generating go files.
.PHONY: pb
pb: cli.install
	@gf gen pb

# Generate protobuf files for database tables.
.PHONY: pbentity
pbentity: cli.install
	@gf gen pbentity