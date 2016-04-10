BASE_PATH = $(PWD)
GOSRC_PATH = $(BASE_PATH)
OUTPUT_PATH = $(BASE_PATH)/build
DIST_PATH = $(BASE_PATH)/dist
ASSETS = $(OUTPUT_PATH)/public/... $(OUTPUT_PATH)/templates/...

LDFLAGS = -ldflags '-X main.BuildSHA=$(BUILD_SHA) -X main.AppName=$(APP_NAME) -X main.Version=$(VERSION)'
DOCKER_TAG = $(OUTPUT_NAME):latest
DOCKER_IMAGE = $(OUTPUT_NAME).tar

BINDATA_BIN = go-bindata -debug
DOCKER_COMPOSE_BIN = docker-compose
GIT_BIN = git
GOBUILD = go build
GULP_BIN = ./node_modules/.bin/gulp
NODE_BIN = node
GULP_OPTIONS =

OUTPUT_NAME = app
APP_NAME = $(shell $(NODE_BIN) -e "console.log(require('./package.json')['name'])")
VERSION = $(shell $(NODE_BIN) -e "console.log(require('./package.json')['version'])")
BUILD_SHA = $(shell $(GIT_BIN) rev-parse HEAD)

ifdef DOCKER
GOBUILD = GOOS=linux GOARCH=386 go build -a -tags netgo -installsuffix netgo
PRODUCTION = 1
endif

ifdef PRODUCTION
BINDATA_BIN = go-bindata
GULP_OPTIONS = --production
endif

all: $(DIST_PATH)/$(OUTPUT_NAME)

.PHONY: clean
clean:
	rm -rf $(DIST_PATH)
	rm -rf $(OUTPUT_PATH)
	rm $(GOSRC_PATH)/assets.go

.PHONY: assets
assets:
	$(GULP_BIN) $(GULP_OPTIONS)

$(OUTPUT_PATH)/public/favicon.ico:
	$(GULP_BIN) $(GULP_OPTIONS)

.PHONY: debends
debends:
	govend -luv

$(GOSRC_PATH)/assets.go: $(OUTPUT_PATH)/public/favicon.ico
	$(BINDATA_BIN) -o $(GOSRC_PATH)/assets.go -prefix build $(ASSETS)

.PHONY: app
app: $(GOSRC_PATH)/assets.go
	$(GOBUILD) $(LDFLAGS) -v -o $(DIST_PATH)/$(OUTPUT_NAME) .

$(DIST_PATH)/$(OUTPUT_NAME): $(GOSRC_PATH)/assets.go
	$(GOBUILD) $(LDFLAGS) -o $(DIST_PATH)/$(OUTPUT_NAME) .

# Docker items

.PHONY: build_docker_server
build_docker_server: $(DIST_PATH)/$(OUTPUT_NAME)
	$(DOCKER_COMPOSE_BIN) build --no-cache

.PHONY: docker_server
docker_server:
	$(DOCKER_COMPOSE_BIN) up -d

.PHONY: stop_docker_server
stop_docker_server:
	$(DOCKER_COMPOSE_BIN) down

.PHONY: rm_docker_server
rm_docker_server:
	$(DOCKER_COMPOSE_BIN) down --rmi all
