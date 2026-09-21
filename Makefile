GOOS=linux
GOARCH=arm64

build-auth:
	mkdir -p dist/auth
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/auth/bootstrap ./cmd/auth
	cd dist/auth && zip -FS function.zip bootstrap

build-planner:
	mkdir -p dist/planner
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/planner/bootstrap ./cmd/planner
	cd dist/planner && zip -FS function.zip bootstrap

build-account:
	mkdir -p dist/account
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/account/bootstrap ./cmd/account
	cd dist/account && zip -FS function.zip bootstrap


build: build-auth build-planner build-account

clean:
	rm -rf dist