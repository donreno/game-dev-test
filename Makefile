BUILDPATH=$(CURDIR)
BINARY=game-dev-test

makedir:
	@echo "Making dirs 📁"
	@if [ ! -d $(BUILDPATH)/build ] ; then mkdir -p $(BUILDPATH)/build ; fi
	@if [ ! -d $(BUILDPATH)/build/sprites ] ; then mkdir -p $(BUILDPATH)/build/sprites ; fi

mod:
	@echo "Vendoring... 📦"
	@go mod vendor

copyassets:
	@echo "Copying assets 🎱"
	@cp sprites/* build/sprites/

build: clean makedir copyassets
	@echo "Building... 🔧"
	@go build -mod vendor -ldflags "-s -w" -o $(BUILDPATH)/build/${BINARY} main.go
	@echo "Binario generado en build/bin/"${BINARY}

test:
	@echo "Running tests... 🧐"
	@go test ./... -v

coverage:
	@echo "Getting coverage... ☔"
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

clean:
	@echo "Cleaning up... 🧹"
	@if [ -d $(BUILDPATH)/build/bin ] ; then rm -rf $(BUILDPATH)/build/ ; fi
	@rm -rf coverfile_out*

benchmark:
	@echo "Running benchmarks... 😎"
	@go test -benchmem ./... -bench=.

.PHONY: clean build test
