all: directories build/api build/webui

clean:
	rm -rf build
	rm -rf webui/etc

run: all
	caddy

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
