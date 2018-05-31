all: directories build/api build/webui

clean:
	rm -rf build
	rm -rf webui/etc

run: all
	caddy

docker-image: all
	# Create a temporary directory.
	$(eval TEMP_DIR := $(shell mktemp -d))

	# Download the caddy package.
	curl -o $(TEMP_DIR)/caddy.tar.gz \
		"https://caddyserver.com/download/linux/amd64?license=personal&telemetry=off"

	# Add application files to the workspace.
	cp -r build $(TEMP_DIR)/files

	# Add Caddyfile to the workspace.
	cp Caddyfile.docker $(TEMP_DIR)/files/Caddyfile

	# Add the caddy binary to the workspace.
	tar xzf $(TEMP_DIR)/caddy.tar.gz -C $(TEMP_DIR)/files caddy

	# Add the Dockerfile outside the workspace.
	cp Dockerfile $(TEMP_DIR)

	# Build the image.
	docker build --rm -t torest $(TEMP_DIR)

	# Remove all used temporary files.
	rm -r $(TEMP_DIR)

test-api:
	go test -cover -v ./api/...

build/api: $(shell ls api/*.go)
	go build -o build/api -v ./api/...

build/webui: webui/node_modules $(shell ls webui/src/**/*) $(shell ls webui/public/index.html) $(shell ls webui/*.*)
	npm --prefix webui run build
	rm -rf webui/etc
	rm -rf build/webui
	mv webui/build build/webui

webui/node_modules:
	npm --prefix webui install
	rmdir -rf webui/etc

directories: build

build:
	mkdir -p build
