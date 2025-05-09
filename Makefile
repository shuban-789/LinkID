GOOS=linux
GOARCH=amd64

all: build

build:
	@echo "Compiling src/*.go into ./linkid"
	GOARCH=$(ARCH) GOOS=$(OS) go build -o linkid ./src

binclean:
	@echo "Removing ./linkid"
	rm linkid

clean:
	@echo "Removing current records from /records"
	rm records/*