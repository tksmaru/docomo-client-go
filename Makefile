
GO_CMD = go
GO_APP_CMD = goapp

# invoke unit test
.PHONY: test
test:
	$(GO_CMD) test -v -cover -coverprofile=coverage.out

# show coverage report on your browser
.PHONY: show_cover
show_cover: test
	sed -i -e "s#.*/\(.*\.go\)#\./\\1#" coverage.out
	$(GO_APP_CMD) tool cover -html=coverage.out

# remove coverage report files
.PHONY: clean
clean:
	-rm -f coverage.out*
