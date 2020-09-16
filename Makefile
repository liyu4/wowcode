 # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=mybinary
    BINARY_UNIX=$(BINARY_NAME)_unix
    test:
		$(GOTEST) -count=1 -v ./...
