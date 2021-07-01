GOBIN=`which go`
EXEC=modmulcircle
$(EXEC):
	$(GOBIN) get && $(GOBIN) build -o $(EXEC)
clean:
	rm -v $(EXEC)
install:
	cp -v $(EXEC) /usr/local/bin/
