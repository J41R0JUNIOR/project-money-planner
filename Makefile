GOOS=linux
GOARCH=arm64

build-auth:
	mkdir -p dist/auth
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/auth/bootstrap ./cmd/auth
	cd dist/auth && zip -FS function.zip bootstrap

build: build-auth

clean:
	rm -rf dist