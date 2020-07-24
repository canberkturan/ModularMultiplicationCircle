GOBIN = /usr/bin/go build
EXEC=modmulcircle
$(EXEC):
	$(GOBIN) -o $(EXEC)
clean:
	rm -v $(EXEC)
install:
	cp -v $(EXEC) /usr/local/bin/
