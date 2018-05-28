all: directories build/torest build/webui

clean:
	rm -r build

directories: build

build:
	mkdir -p build

build/torest: $(shell ls api/*.go)
	go build -o build/torest -v ./api/...

build/webui: $(shell ls webui/src/**/*) $(shell ls webui/public/index.html) $(shell ls webui/*.*)
	npm --prefix ./webui/ run build
	mv webui/build build/webui

test-api:
	go test -cover -v ./api/...
