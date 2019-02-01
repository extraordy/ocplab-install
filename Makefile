GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test ./...
BINFILE=ocplab-install

build:
	$(GOBUILD) -o $(BINFILE) -v

clean:
	$(GOCLEAN)
	rm -f $(BINFILE)

install:
	mv $(BINFILE) /usr/local/bin
