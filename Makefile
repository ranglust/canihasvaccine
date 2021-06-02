MY_VERSION := $(shell git tag | tail -1)

all: clean windows darwin linux

version:
	@echo $(MY_VERSION)

windows:
	@echo "*** building for windows ***"
	GOOS=windows go build -ldflags "-s -w -X github.com/ranglust/canihasvaccine/cmd.VERSION=$(MY_VERSION)" -o "target/canihasvaccine.exe"
	@chmod +x target/canihasvaccine.exe

darwin:
	@echo "*** building for darwin ***"
	GOOS=darwin go build -ldflags "-s -w -X github.com/ranglust/canihasvaccine/cmd.VERSION=$(MY_VERSION)" -o "target/canihasvaccine-darwin"
	@chmod +x target/canihasvaccine-darwin

linux:
	@echo "*** building for darwin ***"
	GOOS=darwin go build -ldflags "-s -w -X github.com/ranglust/canihasvaccine/cmd.VERSION=$(MY_VERSION)" -o "target/canihasvaccine-linux"
	@chmod +x target/canihasvaccine-darwin

clean:
	@echo "removing target"
	@rm -rf target
