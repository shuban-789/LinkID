GOOS=linux
GOARCH=amd64

all: build

build:
	@echo "Compiling src/blockchain.go into ./linkid"
	GOARCH=$(ARCH) GOOS=$(OS) go build -o linkid src/linkid.go

run:
	@echo "Running src/linkid.go"
	go run src/linkid.go

clean:
	@echo "Removing current records from /records"
	rm records/*