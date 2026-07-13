GOOS=linux
GOARCH=arm64

build-auth:
	mkdir -p dist/auth
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/auth/bootstrap ./cmd/auth
	cd dist/auth && zip -FS function.zip bootstrap

build-user:
	mkdir -p dist/user
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/user/bootstrap ./cmd/user
	cd dist/user && zip -FS function.zip bootstrap

build: build-auth build-user

clean:
	rm -rf dist