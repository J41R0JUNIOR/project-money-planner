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

build-transaction:
	mkdir -p dist/transaction
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o dist/transaction/bootstrap ./cmd/transaction
	cd dist/transaction && zip -FS function.zip bootstrap

build: build-auth build-user build-transaction

clean:
	rm -rf dist